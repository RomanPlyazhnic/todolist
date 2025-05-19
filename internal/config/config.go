package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Data struct {
	Name     string        `yaml:"name" env:"NAME" env-default:"todolist"`
	Protocol string        `yaml:"protocol" env:"PROTOCOL" env-default:"rest"` // "rest" or "grpc"
	Port     int           `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout  time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"10s"`
	Env      string        `yaml:"env" env:"ENV" env-default:"todolist"`
	Version  string        `yaml:"version" env:"VERSION" env-default:"0.0.1"`
}

func FromPath(path string) *Data {
	const op = "config.New"

	var cfg Data

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	return &cfg
}

func FromEnv() *Data {
	const op = "config.New"

	var cfg Data

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	return &cfg
}
