package cmd

import (
	"BijanRegmi/envNinja/ent"
	"BijanRegmi/envNinja/pkg"
	"BijanRegmi/envNinja/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	queryAll bool

	queryCmd = &cobra.Command{
		Use:   "query",
		Short: "Query one or all secret's value",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if queryAll {
				secrets, err := envninja.QueryAll(projectName, projectScope)
				if err != nil {
					return err
				}
				utils.PrintSecrets(secrets)
				return nil
			} else if len(args) < 1 {
				return fmt.Errorf("No key provided to query")
			} else {
				secret, err := envninja.Query(projectName, projectScope, args[0])
				if err != nil {
					return err
				}
				utils.PrintSecrets([]*ent.Secret{secret})
				return nil
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(queryCmd)

	queryCmd.Flags().BoolVar(&queryAll, "all", false, "Query all secrets")
}
