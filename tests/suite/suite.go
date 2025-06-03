package suite

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

type Suite struct {
	T   *testing.T
	Ctx context.Context
	App *server.App
}

func New(t *testing.T) *Suite {
	t.Helper()
	t.Parallel()

	cfg, err := setupConfig()
	if err != nil {
		t.Fatal(err)
	}

	a := app.Build(cfg)

	go func() {
		a.Run()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		t.Helper()
		cancel()
		a.Shutdown()
	})

	return &Suite{
		T:   t,
		App: a,
		Ctx: ctx,
	}
}

// setupConfig takes .yml config if --config option is provided
// Otherwise - configuration from ENV variables
func setupConfig() (*config.Data, error) {
	//var cfgPath string
	//flag.StringVar(&cfgPath, "config", "", "config path")
	//flag.Parse()
	path := filepath.Join("config", "test-server.yml")

	return config.FromPath(path)
}
