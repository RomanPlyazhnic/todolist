// HTTP server implementation

package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"net/http"
	"strconv"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/app/server/rest/handlers"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Server represents the main application object
type RestServer struct {
	srv    *http.Server
	router *chi.Mux
}

// New initializes and setups the application
func New(app *server.App, cfg *config.Data) *RestServer {
	router := chi.NewRouter()

	httpSrv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	srv := RestServer{
		srv:    httpSrv,
		router: router,
	}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(httplog.RequestLogger(app.Logger))
	router.Use(middleware.Recoverer)
	if cfg.JWT.Enabled {
		router.Use(JWTAuth(app))
	}

	srv.handleRoutes(app)

	return &srv
}

// Start launches the server
func (s *RestServer) Start(app *server.App) {
	const op = "server.start"

	app.Logger.Info("starting server")
	err := s.srv.ListenAndServe()
	if err != http.ErrServerClosed {
		app.Logger.Error("%s: %v", op, err)
		panic(err)
	}
	app.Logger.Info("app closed", op, s.srv.ListenAndServe())
}

// Shutdown stops the server
func (s *RestServer) Shutdown(app *server.App) {
	const op = "server.shutdown"

	app.Logger.Info("shutting down...", op, true)
	if err := s.srv.Shutdown(context.Background()); err != nil {
		app.Logger.Error("%s: %v", op, err)
	}
}

// handleRoutes describes application's routes
func (s *RestServer) handleRoutes(app *server.App) {
	s.router.Get("/Health", handlers.Health(app))
	s.router.Get("/", handlers.Root(app))
	s.router.Post("/Login", handlers.Login(app))
	s.router.Post("/Register", handlers.Register(app))
}
