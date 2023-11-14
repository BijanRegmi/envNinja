package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetDefaultProjectName() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return filepath.Base(dir)
}

func GetDefaultScope() string {
	return "dev"
}
