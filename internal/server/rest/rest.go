package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RomanPlyazhnic/todolist/internal/config"
)

type Srv struct {
	srv    *http.Server
	router *chi.Mux
	Config *config.Data
	Logger *httplog.Logger
}

func New(cfg *config.Data) *Srv {
	srv := server(cfg)

	srv.router.Use(middleware.RequestID)
	srv.router.Use(middleware.RealIP)
	srv.router.Use(httplog.RequestLogger(srv.Logger))
	srv.router.Use(middleware.Recoverer)
	srv.handleRoutes()

	return srv
}

func (s *Srv) Start() {
	const op = "server.stop"

	s.Logger.Info("server closed", op, s.srv.ListenAndServe())
}

func (s *Srv) Shutdown() {
	const op = "server.shutdown"

	s.Logger.Info("shutting down...", op, true)
	if err := s.srv.Shutdown(context.Background()); err != nil {
		s.Logger.Error("%s: %v", op, err)
	}
}

func server(cfg *config.Data) *Srv {
	r := chi.NewRouter()
	logger := logger(cfg)

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	return &Srv{
		srv:    srv,
		router: r,
		Config: cfg,
		Logger: logger,
	}
}

func logger(cfg *config.Data) *httplog.Logger {
	return httplog.NewLogger(cfg.Name, httplog.Options{
		LogLevel:       slog.LevelDebug,
		Concise:        true,
		RequestHeaders: true,
		Tags: map[string]string{
			"version": cfg.Version,
			"env":     cfg.Env,
		},
	})
}

func (s *Srv) handleRoutes() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
}
