// Test suite for functi9onal tests
// Each test case is running in an isolated database, which is cleared after the test is finished

package suite

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

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
	// TODO: make tests parallel
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
	const op = "suite.waitServerStarted"

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

func configPath() string {
	var path string
	path = os.Getenv("CONFIG_PATH")

	if path == "" {
		path = filepath.Join("config", "test-server.yml")
	}

	return path
}
