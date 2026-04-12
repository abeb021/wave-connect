package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AuthServiceURL    string
	ChatServiceURL    string
	FeedServiceURL    string
	ProfileServiceURL string
	JWTSecret         string
}

func Load() *Config {
	v := viper.New()

	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	// Set defaults
	// v.SetDefault("DB_HOST", "db")
	// v.SetDefault("DB_PORT", "5432")
	// v.SetDefault("DB_NAME", "chat_db")
	// v.SetDefault("DB_USER", "chat_user")
	// v.SetDefault("DB_PASSWORD", "password")
	// v.SetDefault("DB_SSLMODE", "disable")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config: %v", err)
		}
	}

	cfg := &Config{
		AuthServiceURL:    v.GetString("AUTH_SERVICE_URL"),
		ChatServiceURL:    v.GetString("CHAT_SERVICE_URL"),
		FeedServiceURL:    v.GetString("FEED_SERVICE_URL"),
		ProfileServiceURL: v.GetString("PROFILE_SERVICE_URL"),
		JWTSecret:         v.GetString("JWT_SECRET"),
	}

	return cfg
}
