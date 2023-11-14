package envninja

import (
	"BijanRegmi/envNinja/ent"
	"BijanRegmi/envNinja/ent/project"
	"context"
	"fmt"
)

func ListProjects() ([]*ent.Project, error) {
	ctx := context.Background()
	projects, err := DB.Project.Query().All(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed creating project: %w", err)
	}

	return projects, nil
}

func ListScopes(projectName string) ([]*ent.Project, error) {
	ctx := context.Background()
	projects, err := DB.Project.Query().Where(project.Name(projectName)).All(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed creating project: %w", err)
	}

	return projects, nil
}
