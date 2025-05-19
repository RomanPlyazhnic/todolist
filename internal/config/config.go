package config

import "time"

type Data struct {
	Name     string
	Protocol string // "rest" or "grpc"
	Port     int
	Timeout  time.Duration
	Env      string
	Version  string
}

func New() *Data {
	return &Data{
		Name:     "TODOList",
		Protocol: "rest",
		Port:     8080,
		Timeout:  10 * time.Second,
		Env:      "dev",
		Version:  "v0.0.1",
	}
}
