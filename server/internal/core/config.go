package core

import "os"

type Config struct {
	Port          string
	AllowedOrigin string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	origin := os.Getenv("ALLOWED_ORIGIN")

	return &Config{
		Port:          port,
		AllowedOrigin: origin,
	}
}
