/*
Copyright © 2023 candbright <2685082823@qq.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var applicationName string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init the project",
	Long:  `init the project`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(initializeProject(args))
	},
}

func init() {
	initCmd.Flags().StringVarP(&applicationName, "name", "n", "helloworld", "your application name")
	rootCmd.AddCommand(initCmd)
}

func initializeProject(args []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	p := &Project{PkgName: "github.com/candbright/go-core", AbsolutePath: wd, ApplicationName: applicationName}
	return p.Create()
}
