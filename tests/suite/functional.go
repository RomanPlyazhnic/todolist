// Test suite for functional tests
// Each test case is running in an isolated database, which is cleared after the test is finished

package suite

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"testing"
	"time"

	"github.com/RomanPlyazhnic/todolist/internal/app"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/config"
)

func NewFunctional(t *testing.T) *Suite {
	t.Helper()
	// TODO: run only one server instance

	cfg, err := config.Get(configPath())
	if err != nil {
		t.Fatal(err)
	}

	a := app.Build(cfg)

	go func() {
		a.Run()
	}()

	waitServerStarted(a)

	ctx, cancel := context.WithCancel(context.Background())

	t.Cleanup(func() {
		t.Helper()
		cancel()
		a.Logger.Info("test finished", "test", t.Name())
		a.Shutdown()
	})

	return &Suite{
		T:   t,
		App: a,
		Ctx: ctx,
	}
}

func waitServerStarted(a *server.App) {
	const op = "suite.functional.waitServerStarted"

	url := fmt.Sprintf("http://%s:%d/Health", a.Config.Domain, a.Config.Port)
	for i := 0; i < 50; i++ {
		resp, err := http.Get(url)
		if err != nil {
			a.Logger.Error("failed to get health check", op, err)
			continue
		}

		err = resp.Body.Close()
		if err != nil {
			a.Logger.Error("failed to close response body", op, err)
		}
		if resp.StatusCode == http.StatusOK {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}
