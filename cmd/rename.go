package cmd

import (
	envninja "BijanRegmi/envNinja/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	renameKey string

	renameCmd = &cobra.Command{
		Use:   "rename NEW_NAME",
		Short: "Rename secret's key or scope or project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			newName := args[0]

			if cmd.Flags().Changed("key") {
				return envninja.RenameKeyName(projectName, projectScope, renameKey, newName)
			} else if cmd.Flags().Changed("scope") {
				return envninja.RenameScopeName(projectName, projectScope, newName)
			} else if cmd.Flags().Changed("name") {
				return envninja.RenameProjectName(projectName, newName)
			} else {
				return fmt.Errorf("Provide either key or scope or project to remove.")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(renameCmd)

	renameCmd.Flags().StringVarP(&renameKey, "key", "k", "", "secret's key")
}
