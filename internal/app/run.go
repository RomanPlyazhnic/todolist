package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/app/server/rest"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Run launches the server
// Stops when Ctrl+C is pressed
func Run(cfg *config.Data) {
	var srv server.Server

	switch cfg.Protocol {
	case "rest":
		srv = rest.New(cfg)
	// TODO: implement gRPC
	default:
		srv = rest.New(cfg)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		srv.Start()
	}()

	<-done
	srv.Shutdown()
}
