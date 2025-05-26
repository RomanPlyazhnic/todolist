// Build and run the server

package app

import (
	"github.com/RomanPlyazhnic/todolist/internal/app/database"
	"github.com/go-chi/httplog/v2"
	"log/slog"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/app/server/rest"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Build builds the application
// Returns the application object
// cfg - application configuration
func Build(cfg *config.Data) *server.App {
	app := server.App{}

	app.SetConfig(cfg)
	buildLogger(&app, cfg)
	buildServer(&app, cfg)
	buildDatabase(&app, cfg)

	return &app
}

// buildLogger builds the logger
// cfg - application configuration
// app - application object
func buildLogger(app *server.App, cfg *config.Data) {
	logger := httplog.NewLogger(cfg.Name, httplog.Options{
		LogLevel:       slog.LevelDebug,
		Concise:        true,
		RequestHeaders: true,
		Tags: map[string]string{
			"version": cfg.Version,
			"env":     cfg.Env,
		},
	})

	app.SetLogger(logger)
}

// buildServer builds the server
// cfg - application configuration
// app - application object
func buildServer(app *server.App, cfg *config.Data) {
	switch cfg.Protocol {
	case "rest":
		app.SetServer(rest.New(app, cfg))
	// TODO: implement gRPC
	default:
		app.SetServer(rest.New(app, cfg))
	}
}

// buildDatabase builds the database
// cfg - application configuration
// app - application object
func buildDatabase(app *server.App, cfg *config.Data) {
	app.SetDB(database.NewSqliteDB())
}
