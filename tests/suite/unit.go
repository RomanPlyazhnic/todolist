package suite

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"testing"

	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

func NewUnit(t *testing.T) *Suite {
	t.Helper()
	t.Parallel()

	cfg, err := config.Get(configPath())
	if err != nil {
		t.Fatal(err)
	}

	a := app.Build(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	t.Cleanup(func() {
		t.Helper()
		cancel()
		a.Logger.Info("test finished", "test", t.Name())
	})

	return &Suite{
		T:   t,
		App: a,
		Ctx: ctx,
	}
}
