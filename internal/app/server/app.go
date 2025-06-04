// Build and run the server

package server

import (
	"database/sql"
	"github.com/go-chi/httplog/v2"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// App represents the main application object
type App struct {
	srv    Server
	done   chan os.Signal
	Config *config.Data
	Logger *httplog.Logger
	DB     DB
}

// Server represents the server interface
type Server interface {
	Start(*App)
	Shutdown(*App)
}

type DB interface {
	Start(*App) error
	Stop(*App) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Begin() (*sql.Tx, error)
}

// Run launches the server
// Stops when Ctrl+C is pressed
func (a *App) Run() {
	const op = "app.run"

	a.done = make(chan os.Signal, 1)
	signal.Notify(a.done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := a.DB.Start(a)
		if err != nil {
			a.Logger.Error("failed to start database", op, err)
			panic(err)
		}
		a.srv.Start(a)
	}()

	defer func() {
		err := a.DB.Stop(a)
		if err != nil {
			a.Logger.Error("failed to stop database", op, err)
		}
		a.srv.Shutdown(a)
	}()

	<-a.done
}

func (a *App) Shutdown() {
	const op = "app.shutdown"

	a.Logger.Info("shutting down", op, true)

	a.done <- os.Interrupt
}

func (a *App) SetServer(srv Server) {
	a.srv = srv
}

func (a *App) SetConfig(config *config.Data) {
	a.Config = config
}

func (a *App) SetLogger(logger *httplog.Logger) {
	a.Logger = logger
}

func (a *App) SetDB(db DB) {
	a.DB = db
}
