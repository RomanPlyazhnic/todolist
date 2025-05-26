// SQLite database implementation
//
// Example:
// 	db := SqliteDB{}
// 	db.Start(app)
// 	defer db.Stop(app)
// 	res, err := db.Exec("create table foo (id integer not null primary key, name text);")

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// SqliteDB represents Sqlite database client
type SqliteDB struct {
	db *sql.DB
}

// Initialize database
func NewSqliteDB() *SqliteDB {
	return new(SqliteDB)
}

// Start opens database
// Returns error if not opened
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

// Stop closes a database and prevents new queries to start
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

// Exec executes a query
func (db *SqliteDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}
