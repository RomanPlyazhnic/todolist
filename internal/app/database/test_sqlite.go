// Test SQL database, which is used like a regular SQLite, but is cleared when the database is stopped

package database

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/mattn/go-sqlite3"
	"path/filepath"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// SqliteDB represents Sqlite database client
type TestSqliteDB struct {
	db *sql.DB
}

// Initialize database
func NewTestSqliteDB() *TestSqliteDB {
	return new(TestSqliteDB)
}

// Start opens database
// Returns error if not opened
func (db *TestSqliteDB) Start(a *server.App) error {
	const op = "database.Setup"

	a.Logger.Info("starting database", op, true)

	dbPath := filepath.Join(a.Config.RootPath, a.Config.Database.Path)
	database := sql.OpenDB(txdb.New("sqlite3", dbPath))
	a.Logger.Info("database started", op, true)

	db.db = database

	return nil
}

// Stop closes a database and prevents new queries to start
func (db *TestSqliteDB) Stop(a *server.App) error {
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
func (db *TestSqliteDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

// Exequte a query with rows result
// NOTE: don't forget to Close() returned rows
func (db *TestSqliteDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

// Execute a query with 1 row result
func (db *TestSqliteDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(query, args...)
}

// Begin starts a transaction
func (db *TestSqliteDB) Begin() (*sql.Tx, error) {
	return db.db.Begin()
}
