// Launch server

package server

import (
	"github.com/go-chi/httplog/v2"

	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Server is simple server abstraction
type Server interface {
	Start()
	Shutdown()
	Config() *config.Data
	Logger() *httplog.Logger
}
