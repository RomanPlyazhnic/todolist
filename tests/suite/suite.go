package suite

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

type Suite struct {
	T   *testing.T
	Ctx context.Context
	App *server.App
}

func configPath() string {
	var path string
	path = os.Getenv("CONFIG_PATH")

	if path == "" {
		path = filepath.Join("config", "test-server.yml")
	}

	return path
}
