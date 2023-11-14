package cmd

import (
	"BijanRegmi/envNinja/pkg"

	"github.com/spf13/cobra"
)

var (
	pushEnvPath string
	pushKey     string
	pushValue   string

	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "Push secrets from file or key/value pair",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().Changed("path") {
				return envninja.PushFromFile(projectName, projectScope, pushEnvPath)
			} else {
				return envninja.PushKeyValue(projectName, projectScope, pushKey, pushValue)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(pushCmd)

	pushCmd.Flags().StringVarP(&pushKey, "key", "k", "", "KEY")
	pushCmd.Flags().StringVarP(&pushValue, "value", "v", "", "VALUE")
	pushCmd.MarkFlagsRequiredTogether("key", "value")

	pushCmd.Flags().StringVarP(&pushEnvPath, "path", "p", ".env", "Env path")

	pushCmd.MarkFlagsMutuallyExclusive("path", "key")
	pushCmd.MarkFlagsMutuallyExclusive("path", "value")
}
