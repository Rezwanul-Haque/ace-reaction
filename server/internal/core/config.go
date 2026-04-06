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
	if origin == "" {
		origin = "http://localhost:5173"
	}

	return &Config{
		Port:          port,
		AllowedOrigin: origin,
	}
}
