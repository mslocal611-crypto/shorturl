package config

import (
	"cmp"
	"flag"
	"os"
)

type Config struct {
	HTTPAddr string
	BaseURL  string
}

func Load(args []string) (Config, error) {
	cfg := Config{
		HTTPAddr: cmp.Or(os.Getenv("SERVER_ADDRESS"), "localhost:8080"),
		BaseURL:  cmp.Or(os.Getenv("BASE_URL"), "http://localhost:8080"),
	}

	fs := flag.NewFlagSet("shortener", flag.ContinueOnError)
	fs.StringVar(&cfg.HTTPAddr, "a", cfg.HTTPAddr, "адрес запуска HTTP-сервера")
	fs.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "базовый URL для короткой ссылки")

	err := fs.Parse(args)

	return cfg, err
}
