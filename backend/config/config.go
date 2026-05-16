package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	Migration MigrationConfig `yaml:"migration"`
}

type MigrationConfig struct {
	Path string `yaml:"path"`
	Enabled bool `yaml:"enabled"`
}

type AppConfig struct {
	Name string     `yaml:"name"`
	Port string     `yaml:"port"`
	Db   DbConfig   `yaml:"db"`
	Cors CorsConfig `yaml:"cors"`
}

type CorsConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedHeaders []string `yaml:"allowed_headers"`
	AllowCredentials bool    `yaml:"allow_credentials"`
}

type DbConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	DBHost     string `yaml:"dbhost"`
	DBUser     string `yaml:"dbuser"`
	DBPassword string `yaml:"dbpassword"`
	DBName     string `yaml:"dbname"`
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName,
	)
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfg := &Config{}
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return cfg, nil
}
