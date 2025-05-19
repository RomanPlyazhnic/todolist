package main

import (
	"flag"

	"github.com/RomanPlyazhnic/todolist/internal/config"
	"github.com/RomanPlyazhnic/todolist/internal/server"
)

func main() {
	cfg := setupConfig()

	server.Run(cfg)
}

func setupConfig() *config.Data {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "", "config path")
	flag.Parse()

	var cfg *config.Data
	if cfgPath != "" {
		cfg = config.FromPath(cfgPath)
	} else {
		cfg = config.FromEnv()
	}

	return cfg
}
