package cmd

import (
	envninja "BijanRegmi/envNinja/pkg"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove [KEY1 KEY2 .. KEYN]",
		Short: "Remove secrets or scope or project",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return envninja.RemoveSecrets(projectName, projectScope, args)
			} else if cmd.Flags().Changed("scope") {
				return envninja.RemoveScope(projectName, projectScope)
			} else if cmd.Flags().Changed("name") {
				return envninja.RemoveProject(projectName)
			} else {
				return fmt.Errorf("Provide either key to remove or scope or project.")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
