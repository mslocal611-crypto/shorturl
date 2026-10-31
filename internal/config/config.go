package config

import (
	"cmp"
	"os"
)

type Config struct {
	Host string
	Port string
}

func Load() Config {
	return Config{
		Host: cmp.Or(os.Getenv("HOST"), "localhost"),
		Port: cmp.Or(os.Getenv("PORT"), "8080"),
	}
}
