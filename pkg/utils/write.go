package utils

import (
	"bytes"
	"fmt"
	"os"
)

func WriteEnvs(envs map[string]string, path string) error {
	var buffer bytes.Buffer
	for key, value := range envs {
		_, err := buffer.WriteString(fmt.Sprintf("%v = \"%v\"\n", key, value))
		if err != nil {
			return err
		}
	}
	return os.WriteFile(path, buffer.Bytes(), 0644)
}
