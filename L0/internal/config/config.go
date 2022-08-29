package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port      string `env:"PORT" envDefault:"8080"`
	ProxyPort string `env:"HTTP_PORT" envDefault:"8081"`
	Subj      string `env:"SUBJECT" envDefault:"foo"`
	DBConn    string `env:"DBConn" envDefault:"host=localhost user=postgres password=postgres dbname=backend port=5432 sslmode=disable"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
