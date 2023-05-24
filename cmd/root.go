/*
Copyright © 2023 candbright <2685082823@qq.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "cli for init",
	Long:  `cli for init`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
