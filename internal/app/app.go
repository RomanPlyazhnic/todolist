// Launch server

package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanPlyazhnic/todolist/internal/app/rest"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Server is simple server abstraction
type Server interface {
	Start()
	Shutdown()
}

// Run launches the server
// Stops when Ctrl+C is pressed
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

// Server returns server implementation
func server(cfg *config.Data) Server {
	switch cfg.Protocol {
	case "rest":
		return rest.New(cfg)
	// TODO: implement gRPC
	default:
		return rest.New(cfg)
	}
}
