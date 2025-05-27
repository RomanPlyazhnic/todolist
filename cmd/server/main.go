// Entry point for the application

package main

import (
	"flag"

	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Launch application
func main() {
	cfg, err := setupConfig()
	if err != nil {
		panic(err)
	}

	a := app.Build(cfg)
	a.Run()
}

// setupConfig takes .yml config if --config option is provided
// Otherwise - configuration from ENV variables
func setupConfig() (*config.Data, error) {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "", "config path")
	flag.Parse()

	return config.FromPath(cfgPath)
}
