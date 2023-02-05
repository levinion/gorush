package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "gorush",
	Short: "Welcome to GoRush!",
	Long:  "GoRush is a Greeeeeat Blog Cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello GoRush!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
