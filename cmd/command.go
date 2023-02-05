package cmd

import "github.com/spf13/cobra"

// 创建静态文件
var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "使用此命令以构建docs目录",
	Run:     Build,
	Aliases: []string{"b"},
}

// 清理公开文件夹
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "使用此命令以清理docs目录，使用'-a'清空所有由系统生成的非内容文件夹[谨慎使用]",
	Run:     Clean,
	Aliases: []string{"c"},
}

// 获取默认模板命令
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "使用此命令以获取默认项目模板，仅用于初始化",
	Run:     Get,
	Aliases: []string{"g"},
}

// 创建新文章命令
var newCmd = &cobra.Command{
	Use:     "new post",
	Short:   "使用此命令以创建新文章，应以dir/file的格式输入参数;使用'-p'创建新页面",
	Run:     New,
	Aliases: []string{"n"},
}

// 创建静态文件并路由
var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "使用此命令以构建docs目录并在本地预览",
	Run:     BuildAndRender,
	Aliases: []string{"s"},
}
