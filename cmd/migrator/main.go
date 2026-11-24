package migrator

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string
	flag.StringVar(&storagePath, "storage-path", "", "path ")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "path")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+storagePath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsPath),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
	}

}
