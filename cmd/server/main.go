// Entry point for the application

package main

import (
	"flag"
	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/config"
	"os"
)

// Launch application
func main() {
	cfg, err := config.Get(configPath())
	if err != nil {
		panic(err)
	}

	a := app.Build(cfg)
	a.Run()
}

func configPath() string {
	var path string
	path = os.Getenv("CONFIG_PATH")

	if path == "" {
		flag.StringVar(&path, "config", "", "config path")
		flag.Parse()
	}

	return path
}
