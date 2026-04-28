package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string

	KafkaBroker string

	JWTSecret string
}

func Load() *Config {
	v := viper.New()

	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	//v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults
	// v.SetDefault("JWT_SECRET", "changeme")
	// v.SetDefault("DB_HOST", "auth_db")
	// v.SetDefault("DB_PORT", "5433")
	// v.SetDefault("DB_NAME", "auth_db")
	// v.SetDefault("DB_USER", "auth_user")
	// v.SetDefault("DB_PASSWORD", "password")
	v.SetDefault("DB_SSLMODE", "disable")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config: %v", err)
		}
	}

	cfg := &Config{
		JWTSecret:   v.GetString("JWT_SECRET"),
		DBHost:      v.GetString("DB_HOST"),
		DBPort:      v.GetString("DB_PORT"),
		DBUser:      v.GetString("DB_USER"),
		DBPassword:  v.GetString("DB_PASSWORD"),
		DBName:      v.GetString("DB_NAME"),
		DBSSLMode:   v.GetString("DB_SSLMODE"),
		KafkaBroker: v.GetString("KAFKA_BROKER"),
	}

	return cfg
}

func (c *Config) DatabaseURL() string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
	return connStr
}
