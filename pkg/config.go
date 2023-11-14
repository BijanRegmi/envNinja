package envninja

import (
	"BijanRegmi/envNinja/ent"
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/lib/pq"
)

var DB *ent.Client

func ConnectDatabase(dsn string) {
	var err error

	DB, err = ent.Open(dialect.SQLite, dsn)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func CloseDatabase() {
	DB.Close()
}
