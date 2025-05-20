// Entry point for the application

package main

import (
	"flag"

	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Launch application
func main() {
	cfg := setupConfig()

	app.Run(cfg)
}

// setupConfig takes .yml config if --config option is provided
// Otherwise - configuration from ENV variables
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
