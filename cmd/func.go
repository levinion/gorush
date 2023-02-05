package cmd

import (
	"gorush/internal/builder"
	"gorush/internal/embed/assets"
	"gorush/internal/embed/config"
	"gorush/internal/embed/root"
	"gorush/internal/embed/templates"
	"gorush/internal/model"
	"gorush/internal/util"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Build(cmd *cobra.Command, args []string) {
	build()
}

func Clean(cmd *cobra.Command, args []string) {
	isCleanAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		zap.L().Error("未定义行为")
	}
	clean(isCleanAll)
}

// 在当前目录下生成默认结构
func Get(cmd *cobra.Command, args []string) {
	getContentDir() //新建content目录结构
	config.Get()    //获取默认配置文件
	templates.Get() //获取默认模板
	assets.Get()    //获取默认css和js文件
	root.Get()
	util.SimpleMkdir("static")    //获取默认静态文件

}

func getContentDir() {
	util.SimpleMkdir("content", "pages")
	util.SimpleMkdir("content", "posts")
	util.SimpleMkdir("content", "posts", "drafts")
}

func New(cmd *cobra.Command, args []string) {
	//判断参数合法性
	if len(args) < 1 {
		zap.L().Error("参数不能为空")
		return
	}
	//获取参数
	addr := args[0]
	isPage, err := cmd.Flags().GetBool("page")
	if err != nil {
		zap.L().Error("读取配置失败，可能不存在该配置项")
	}
	if !isPage {
		//获取元数据
		meta := model.MetaData{
			Title:   util.TrimExt(addr),
			Created: time.Now().Format("2006-01-02 15:04:05"),
		}
		//创建文章模板
		templates.MakeStdPostTemplateFile(addr, meta)
	} else {
		//创建页面模板
		templates.MakeCommonPage(addr)
	}
}

func BuildAndRender(cmd *cobra.Command, args []string) {
	//指定所使用的模板，解析markdown，生成静态文件，路由静态文件
	render(build())
}

func build() *builder.Builder {
	initConfig()
	builder := builder.NewBuilder()
	//暂时使用Basic作为模板，后续在文件中单独指定模板
	clean(false)
	builder.Dump()
	builder.Freeze("Basic")
	zap.L().Info("成功构建docs文件夹")
	return builder
}

func render(builder *builder.Builder) {
	addr := viper.GetString("server.addr")
	builder.Render()
	zap.L().Info("开始监听" + "http://" + addr + "...")
	builder.Run(addr)
}

func clean(isCleanAll bool) {
	var allSystemDir = []string{
		"docs",
		"templates",
		"assets",
		"static",
		"config.toml",
	}
	if isCleanAll {
		for _, dir := range allSystemDir {
			os.RemoveAll(dir)
		}
		zap.L().Info("完成全部项目文件夹清理行为")
	} else {
		os.RemoveAll(filepath.Join("docs"))
		zap.L().Info("完成清理docs文件夹")
	}
}

func initConfig() {
	viper.SetConfigFile("config.toml")
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Debug("配置文件不存在，请尝试运行 gorush get 命令获取默认配置文件")
		os.Exit(1)
	}
}
