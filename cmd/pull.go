package cmd

import (
	envninja "BijanRegmi/envNinja/pkg"

	"github.com/spf13/cobra"
)

var (
	envPath string

	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull secrets to file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return envninja.PullToFile(projectName, projectScope, envPath)
		},
	}
)

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&envPath, "path", "p", ".env", "Env file to pull to")
}
