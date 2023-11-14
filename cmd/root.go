package cmd

import (
	"BijanRegmi/envNinja/pkg/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
	projectName  string
	projectScope string

	rootCmd = &cobra.Command{
		Use:   "envNinja",
		Short: "Manange env like a ninja",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	name := utils.GetDefaultProjectName()
	scope := utils.GetDefaultScope()

	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", name, "Project name")
	rootCmd.PersistentFlags().StringVarP(&projectScope, "scope", "s", scope, "Project scope")
}
