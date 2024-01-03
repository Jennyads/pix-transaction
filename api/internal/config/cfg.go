package config

import "os"

type Config struct {
	Api     ApiConfig
	Profile ProfileConfig
}

type ApiConfig struct {
	Port string
}

type ProfileConfig struct {
	Host string
	Port string
}

func Load() *Config {
	return &Config{
		Profile: ProfileConfig{
			Host: Getenv("PROFILE_BACKEND_HOST", "localhost"),
			Port: Getenv("PROFILE_BACKEND_PORT", "9080"),
		},
		Api: ApiConfig{
			Port: Getenv("API_PORT", "9060"),
		},
	}
}

func Getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
