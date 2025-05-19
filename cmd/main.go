package main

import (
	"github.com/RomanPlyazhnic/todolist/internal/config"
	"github.com/RomanPlyazhnic/todolist/internal/server"
)

func main() {
	cfg := config.New()
	server.Run(cfg)
}
