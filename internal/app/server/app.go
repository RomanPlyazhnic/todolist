// Build and run the server

package server

import (
	"github.com/go-chi/httplog/v2"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// App represents the main application object
type App struct {
	srv    Server
	Config *config.Data
	Logger *httplog.Logger
	//DB     *database.Database
}

// Server represents the server interface
type Server interface {
	Start(*App)
	Shutdown(*App)
}

// Run launches the server
// Stops when Ctrl+C is pressed
func (a *App) Run() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.srv.Start(a)
	}()

	<-done
	a.srv.Shutdown(a)
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
