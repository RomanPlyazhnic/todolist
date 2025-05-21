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
}

// FromPath returns configuration based on .yml file
func FromPath(path string) (*Data, error) {
	const op = "config.New"

	var cfg Data

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return &Data{}, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}

// FromEnv returns configuration based on ENV variables
func FromEnv() (*Data, error) {
	const op = "config.New"

	var cfg Data

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &Data{}, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}
