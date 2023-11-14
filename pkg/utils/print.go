package utils

import (
	"BijanRegmi/envNinja/ent"
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintSecrets(secrets []*ent.Secret) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value"})

	for _, secret := range secrets {
		table.Append([]string{secret.Key, secret.Value})
	}
	table.Render()
}

func PrintProjects(projects []*ent.Project) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project", "Scope"})

	for _, project := range projects {
		table.Append([]string{project.Name, project.Scope})
	}
	table.Render()
}
