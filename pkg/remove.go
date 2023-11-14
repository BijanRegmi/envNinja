package envninja

import (
	"BijanRegmi/envNinja/ent/project"
	"BijanRegmi/envNinja/ent/secret"
	"context"
)

func RemoveProject(name string) error {
	ctx := context.Background()

	_, err := DB.Project.Delete().Where(project.Name(name)).Exec(ctx)

	return err
}

func RemoveScope(name string, scope string) error {
	ctx := context.Background()

	_, err := DB.Project.Delete().Where(project.Name(name), project.Scope(scope)).Exec(ctx)

	return err

}

func RemoveSecrets(name string, scope string, keys []string) error {
	ctx := context.Background()

	_, err := DB.Secret.Delete().Where(secret.KeyIn(keys...), secret.HasProjectWith(project.Name(name), project.Scope(scope))).Exec(ctx)

	return err
}
