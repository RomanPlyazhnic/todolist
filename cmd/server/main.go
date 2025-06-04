// Entry point for the application

package main

import (
	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

// Launch application
func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	a := app.Build(cfg)
	a.Run()
}
