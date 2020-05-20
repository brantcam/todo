package config

import (
	"fmt"
	"os"
)

// Config contains db connection params
type Config struct {
	DBUser string
	DBPass string
	DBName string
	DBPort string
	DBHost string
}

// LoadConfigFromEnv grabs all db connection params from env file
func LoadConfigFromEnv() (*Config, error) {
	var c Config

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, fmt.Errorf("DB_USER is empty")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		return nil, fmt.Errorf("DB_PASS is empty")
	}

	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, fmt.Errorf("DB_NAME is empty")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("DB_PORT is empty")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, fmt.Errorf("DB_HOST is empty")
	}

	c.DBUser = user
	c.DBPass = pass
	c.DBName = name
	c.DBPort = port
	c.DBHost = host

	return &c, nil
}
