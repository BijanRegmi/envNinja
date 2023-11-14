package envninja

import (
	"BijanRegmi/envNinja/ent"
	"BijanRegmi/envNinja/ent/project"
	"BijanRegmi/envNinja/ent/secret"
	"context"
	"fmt"
)

func Query(name string, scope string, key string) (*ent.Secret, error) {
	ctx := context.Background()
	secret, err := DB.Secret.Query().Where(secret.Key(key), secret.HasProjectWith(project.Name(name), project.Scope(scope))).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to query secret: %w", err)
	}

	return secret, nil
}

func QueryAll(name string, scope string) ([]*ent.Secret, error) {
	ctx := context.Background()
	secrets, err := DB.Secret.Query().Where(secret.HasProjectWith(project.Name(name), project.Scope(scope))).All(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to query secrets: %w", err)
	}

	return secrets, nil
}
