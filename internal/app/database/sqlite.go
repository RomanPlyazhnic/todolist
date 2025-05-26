// SQLite database implementation

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

type SqliteDB struct {
	db *sql.DB
}

func NewSqliteDB() *SqliteDB {
	return &SqliteDB{}
}

func (db *SqliteDB) Start(a *server.App) error {
	const op = "database.Setup"

	a.Logger.Info("starting database", op, true)
	database, err := sql.Open("sqlite3", a.Config.Database.Path)
	if err != nil {
		a.Logger.Error("failed to launch database", op, err)
		return fmt.Errorf("failed to launch database", op, err)
	}
	a.Logger.Info("database started", op, true)

	db.db = database

	return nil
}

func (db *SqliteDB) Stop(a *server.App) error {
	const op = "database.Close"

	a.Logger.Info("closing database", op, true)
	err := db.db.Close()
	if err != nil {
		a.Logger.Error("failed to close database", op, err)
		return fmt.Errorf("failed to close database", op, err)
	}
	a.Logger.Info("database closed", op, true)

	return nil
}
