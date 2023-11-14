package cmd

import (
	"BijanRegmi/envNinja/ent"
	envninja "BijanRegmi/envNinja/pkg"
	"BijanRegmi/envNinja/pkg/utils"

	"github.com/spf13/cobra"
)

var (
	listAll bool

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all projects or scopes within a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				projects []*ent.Project
				err      error
			)
			if listAll {
				projects, err = envninja.ListProjects()
			} else {
				projects, err = envninja.ListScopes(projectName)
			}
			if err != nil {
				return err
			}
			utils.PrintProjects(projects)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&listAll, "all", false, "List all projects")
}
