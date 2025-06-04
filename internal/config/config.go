// Load server configuration

package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Data represents configuration format
type Data struct {
	Name     string        `yaml:"name" env:"NAME" env-default:"todolist"`
	Protocol string        `yaml:"protocol" env:"PROTOCOL" env-default:"rest"` // "rest" or "grpc"
	Port     int           `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout  time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"10s"`
	Env      string        `yaml:"env" env:"ENV" env-default:"todolist"`
	Version  string        `yaml:"version" env:"VERSION" env-default:"0.0.1"`
	Domain   string        `yaml:"domain" env:"DOMAIN" env-default:"localhost"`
	RootPath string
	Database Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
}

type Database struct {
	Test bool   `yaml:"test" env:"DATABASE_TEST" env-default:"false"`
	Path string `yaml:"path" env:"DATABASE_PATH" env-default:"./test.db"`
}

type JWT struct {
	Enabled       bool          `yaml:"enabled" env:"JWT_ENABLED" env-default:"false"`
	Secret        string        `env:"JWT_SECRET"`
	TokenDuration time.Duration `yaml:"duration" env:"JWT_DURATION" env-default:"1h"`
}

// Get returns configuration
// If a config path presents - read .yml config and override it with ENV variables
// Config path can be provided via --config=PATH or ENV variable CONFIG_PATH
// If a config path not presents - read config from ENV variables
func Get() (*Data, error) {
	const op = "config.New"

	var cfg Data
	var err error

	var path string
	path = os.Getenv("CONFIG_PATH")

	if path == "" {
		flag.StringVar(&path, "config", "", "config path")
		flag.Parse()
	}

	_, curPath, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(curPath, "..", "..", "..")

	if path != "" {
		err = cleanenv.ReadConfig(filepath.Join(rootPath, path), &cfg)
		if err != nil {
			return &Data{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &Data{}, fmt.Errorf("%s: %w", op, err)
	}

	cfg.RootPath = filepath.Join(curPath, "./../../../")

	return &cfg, nil
}
