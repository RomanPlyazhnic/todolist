// Load server configuration

package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Data represents configuration format
type Data struct {
	Name     string        `yaml:"name" env:"NAME" env-default:"todolist"`
	Protocol string        `yaml:"protocol" env:"PROTOCOL" env-default:"rest"` // "rest" or "grpc"
	Port     int           `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout  time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"10s"`
	Env      string        `yaml:"env" env:"ENV" env-default:"todolist"`
	Version  string        `yaml:"version" env:"VERSION" env-default:"0.0.1"`
	Secret   string        `env:"SECRET"`
}

// FromPath returns configuration
// If config path presents - read .yml config and override it with ENV variables
// If config path not presents - read config from ENV variables
func FromPath(path string) (*Data, error) {
	const op = "config.New"

	var cfg Data
	var err error

	if path != "" {
		err = cleanenv.ReadConfig(path, &cfg)
		if err != nil {
			return &Data{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &Data{}, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}
