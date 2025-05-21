// HTTP server implementation

package rest

import (
	"context"
	"github.com/RomanPlyazhnic/todolist/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"strconv"
)

// App represents the main application object
type App struct {
	srv    *http.Server
	router *chi.Mux
	Config *config.Data
	Logger *httplog.Logger
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

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(httplog.RequestLogger(logger))
	router.Use(middleware.Recoverer)

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	app := App{
		srv:    srv,
		router: router,
		Config: cfg,
		Logger: logger,
	}

	app.handleRoutes()

	return &app
}

// Start launches the server
func (a *App) Start() {
	const op = "app.stop"

	a.Logger.Info("app closed", op, a.srv.ListenAndServe())
}

// Shutdown stops the server
func (a *App) Shutdown() {
	const op = "app.shutdown"

	a.Logger.Info("shutting down...", op, true)
	if err := a.srv.Shutdown(context.Background()); err != nil {
		a.Logger.Error("%s: %v", op, err)
	}
}

// handleRoutes describes application's routes
func (a *App) handleRoutes() {
	a.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
}
