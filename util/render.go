package util

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderHandle(filename string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			basename := filepath.Base(filename)
			t, err := template.New(basename).Delims("{{{", "}}}").ParseFiles(filename)
			if err != nil {
				panic(err)
			}
			t.Execute(w, nil)
		}
	}
}
