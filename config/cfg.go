package config

import "os"

type config struct {
	NamePrefix string
}

func Config() *config {
	return &config{
		NamePrefix: os.Getenv("APP_NAME_PREFIX"),
	}
}
