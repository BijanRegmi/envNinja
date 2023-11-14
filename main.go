package main

import (
	"BijanRegmi/envNinja/cmd"
	"BijanRegmi/envNinja/pkg"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"context"
	"log"
)

func init() {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatalf("Error obtaining current user: %v", err)
			return
		}
		configHome = filepath.Join(currentUser.HomeDir, ".config")
	}

	dbFolderPath := filepath.Join(configHome, "envNinja")

	// Check if the directory doesn't exist
	if _, err := os.Stat(dbFolderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dbFolderPath, 0755); err != nil {
			log.Fatalf("Error creating directory: %v", err)
			return
		}
	}

	// Construct the path to the SQLite database file
	dbFilePath := filepath.Join(dbFolderPath, "envNinja.db")
	dsn := fmt.Sprintf("file:%s?_fk=1", dbFilePath)
	envninja.ConnectDatabase(dsn)
}

func main() {
	defer envninja.CloseDatabase()
	// Run the auto migration tool.
	if err := envninja.DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	cmd.Execute()
}
