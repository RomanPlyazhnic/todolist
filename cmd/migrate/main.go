// Migrate database

package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"path/filepath"

	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Run newest migrations
func main() {
	const op = "database.migrate"

	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	a := app.Build(cfg)
	db, err := sql.Open("sqlite3", filepath.Join(a.Config.RootPath, a.Config.Database.Path))
	if err != nil {
		a.Logger.Error("failed to open database", op, err)
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			a.Logger.Error("failed to close database", op, err)
		}
	}()

	a.Logger.Info("starting database migration", "migrate", true)
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		a.Logger.Error("failed to setup migrate driver", op, err)
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "sqlite3", driver)
	if err != nil {
		a.Logger.Error("failed to setup migrate instance", op, err)
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		a.Logger.Error("failed to migrate", op, err)
	}

	a.Logger.Info("database migrated", op, true)
}
