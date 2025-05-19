package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanPlyazhnic/todolist/internal/config"
	"github.com/RomanPlyazhnic/todolist/internal/server/rest"
)

type Server interface {
	Start()
	Shutdown()
}

func Run(cfg *config.Data) {
	srv := server(cfg)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		srv.Start()
	}()

	<-done
	srv.Shutdown()
}

func server(cfg *config.Data) Server {
	switch cfg.Protocol {
	case "rest":
		return rest.New(cfg)
	// TODO: implement gRPC
	default:
		return rest.New(cfg)
	}
}
