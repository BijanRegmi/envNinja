package envninja

import (
	"BijanRegmi/envNinja/ent"
	"BijanRegmi/envNinja/ent/project"
	"BijanRegmi/envNinja/ent/secret"
	"BijanRegmi/envNinja/pkg/utils"
	"context"
	"io"
	"os"
)

func PushFromFile(name string, scope string, path string) error {
	ctx := context.Background()

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	envs := utils.ParseEnv(data)

	project, err := DB.Project.Query().Where(project.Name(name), project.Scope(scope)).Only(ctx)

	if ent.IsNotFound(err) {
		project, err = DB.Project.
			Create().
			SetName(name).SetScope(scope).
			Save(ctx)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	var secretsCreate []*ent.SecretCreate

	for key, value := range envs {
		secretsCreate = append(secretsCreate, DB.Secret.Create().SetKey(key).SetValue(value).SetProject(project))
	}

	err = DB.Secret.CreateBulk(secretsCreate...).OnConflictColumns(secret.FieldKey, secret.ProjectColumn).UpdateValue().Exec(ctx)
	return err
}

func PushKeyValue(name string, scope string, key string, value string) error {
	ctx := context.Background()

	project, err := DB.Project.Query().Where(project.Name(name), project.Scope(scope)).Only(ctx)

	if ent.IsNotFound(err) {
		project, err = DB.Project.
			Create().
			SetName(name).SetScope(scope).
			Save(ctx)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	err = DB.Secret.Create().SetKey(key).SetValue(value).SetProject(project).OnConflictColumns(secret.FieldKey, secret.ProjectColumn).UpdateValue().Exec(ctx)
	return err
}
