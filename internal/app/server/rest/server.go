// HTTP server implementation

package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RomanPlyazhnic/todolist/internal/app/server/rest/handlers"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// App represents the main application object
type App struct {
	srv    *http.Server
	router *chi.Mux
	config *config.Data
	logger *httplog.Logger
}

// New initializes and setups the application
func New(cfg *config.Data) *App {
	router := chi.NewRouter()

	logger := httplog.NewLogger(cfg.Name, httplog.Options{
		LogLevel:       slog.LevelDebug,
		Concise:        true,
		RequestHeaders: true,
		Tags: map[string]string{
			"version": cfg.Version,
			"env":     cfg.Env,
		},
	})

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	app := App{
		srv:    srv,
		router: router,
		config: cfg,
		logger: logger,
	}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(httplog.RequestLogger(logger))
	router.Use(middleware.Recoverer)
	if cfg.JWT.Enabled {
		router.Use(JWTAuth(&app))
	}

	app.handleRoutes()

	return &app
}

// Start launches the server
func (a *App) Start() {
	const op = "app.stop"

	a.logger.Info("starting server")
	a.logger.Info("app closed", op, a.srv.ListenAndServe())
}

// Shutdown stops the server
func (a *App) Shutdown() {
	const op = "app.shutdown"

	a.logger.Info("shutting down...", op, true)
	if err := a.srv.Shutdown(context.Background()); err != nil {
		a.logger.Error("%s: %v", op, err)
	}
}

func (a *App) Config() *config.Data {
	return a.config
}

func (a *App) Logger() *httplog.Logger {
	return a.logger
}

// handleRoutes describes application's routes
func (a *App) handleRoutes() {
	a.router.Get("/", handlers.Root(a))
	a.router.Post("/Login", handlers.Login(a))
	a.router.Post("/Register", handlers.Register(a))
}
