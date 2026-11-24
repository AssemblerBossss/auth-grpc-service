package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string
	flag.StringVar(&storagePath, "storage-path", "", "path to the SQLite database")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(
		&migrationsTable,
		"migrations-table",
		"migrations",
		"name of the migrations table",
	)
	flag.Parse()

	if storagePath == "" {
		log.Fatal("storage-path is required")
	}

	if migrationsPath == "" {
		log.Fatal("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf(
			"sqlite3://%s?x-migrations-table=%s",
			storagePath,
			url.QueryEscape(migrationsTable),
		),
	)
	if err != nil {
		log.Fatalf("initialize migrator: %v", err)
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("no migrations to apply")
		return
	}
	if err != nil {
		log.Fatalf("apply migrations: %v", err)
	}
	fmt.Println("migrations applied successfully")
}
