package cmd

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(serveCmd)
}

func init() {
	newCmd.Flags().BoolP("page", "p", false, "此参数表明新建页面")
	cleanCmd.Flags().BoolP("all", "a", false, "此参数表明删除所有由系统生成的非内容文件夹")
}
