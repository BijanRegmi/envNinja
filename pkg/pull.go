package envninja

import (
	"BijanRegmi/envNinja/ent/project"
	"BijanRegmi/envNinja/ent/secret"
	"BijanRegmi/envNinja/pkg/utils"
	"context"
)

func PullToFile(name string, scope string, path string) error {
	ctx := context.Background()

	secrets, err := DB.Secret.Query().Where(secret.HasProjectWith(project.Name(name), project.Scope(scope))).All(ctx)
	if err != nil {
		return err
	}

	envs := make(map[string]string, 0)
	for _, secrets := range secrets {
		envs[secrets.Key] = secrets.Value
	}
	return utils.WriteEnvs(envs, path)
}
