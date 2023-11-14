package envninja

import (
	"BijanRegmi/envNinja/ent/project"
	"BijanRegmi/envNinja/ent/secret"
	"context"
)

func RenameProjectName(currentProjectName string, newProjectName string) error {
	ctx := context.Background()

	return DB.Project.Update().Where(project.Name(currentProjectName)).SetName(newProjectName).Exec(ctx)
}

func RenameScopeName(currentProjectName string, currentScopeName string, newScopeName string) error {
	ctx := context.Background()

	return DB.Project.Update().Where(project.Name(currentProjectName), project.Scope(currentScopeName)).SetScope(newScopeName).Exec(ctx)
}

func RenameKeyName(currentProjectName string, currentScopeName string, currentKeyName string, newKeyName string) error {
	ctx := context.Background()

	return DB.Secret.Update().Where(secret.Key(currentKeyName), secret.HasProjectWith(project.Name(currentProjectName), project.Scope(currentScopeName))).SetKey(newKeyName).Exec(ctx)
}
