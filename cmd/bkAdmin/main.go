package main

import (
	"back-admin/cmd/app"
	"back-admin/cmd/setup"
	"back-admin/cmd/version"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// @title       backend-admin
// @version     v1
// @description 后台管理系统
func main() {
	var rootCmd = &cobra.Command{
		Use:   "bkAdmin",
		Short: "bkAdmin",
	}

	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(setup.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cmd execute err:", err.Error())
		os.Exit(-1)
	}
}
