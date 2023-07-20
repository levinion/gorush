package rebirth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type Birth struct {
	f      func(c Context)
	onFunc func()
	*Context
}

type Context struct {
	count  int
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Context) Check() bool {
	select {
	case <-c.ctx.Done():
		return true
	default:
		return false
	}
}

func New() *Birth {
	ctx, cancel := context.WithCancel(context.Background())
	return &Birth{
		Context: &Context{
			ctx:    ctx,
			cancel: cancel,
		},
	}
}

func (b *Birth) Call(f func(c Context)) *Birth {
	b.f = f
	return b
}

func (b *Birth) Cancel() {
	b.cancel()
}

func (b *Birth) Done() {
	go func(c Context) {
		b.count++
		if b.count != 1 {
			fmt.Println("Restarting...")
		}
		b.f(c)
	}(*b.Context)
	if b.onFunc != nil {
		b.onFunc()
	}
}

func (b *Birth) Re() {
	b.Cancel()
	ctx, cancel := context.WithCancel(context.Background())
	b.ctx = ctx
	b.cancel = cancel
	b.Done()
}

func (b *Birth) OnDirChange(dirname string, d time.Duration) *Birth {
	b.onFunc = func() {
		md := checkDirMd5(dirname)
		ticker := time.NewTicker(d)
		for range ticker.C {
			newMd := checkDirMd5(dirname)
			if md != newMd {
				md = newMd
				b.Re()
			}
		}
	}
	return b
}

func (b *Birth) OnFileChange(filename string, d time.Duration) *Birth {
	b.onFunc = func() {
		md := checkMd5(filename)
		ticker := time.NewTicker(d)
		for range ticker.C {
			newMd := checkMd5(filename)
			if md != newMd {
				md = newMd
				b.Re()
			}
		}
	}
	return b
}

func checkMd5(filename string) string {
	file, _ := os.ReadFile(filename)
	md := md5.Sum(file)
	return hex.EncodeToString(md[:])
}

func checkDirMd5(dirname string) string {
	m := &sync.Map{}
	var wg sync.WaitGroup
	filepath.Walk(dirname, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go func() {
			m.Store(path, checkMd5(path))
			wg.Done()
		}()
		return nil
	})
	wg.Wait()
	ks := make([]string, 0)
	m.Range(func(k, value any) bool {
		ks = append(ks, k.(string))
		return true
	})
	sort.Strings(ks)
	bf := ""
	for _, v := range ks {
		v, _ := m.Load(v)
		bf += v.(string)
	}
	r := md5.Sum([]byte(bf))
	rs := hex.EncodeToString(r[:])
	return rs
}
