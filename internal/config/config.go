package config

import "os"

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	Port        string
	FilesAPIURL string
}

func Load() *Config {
	return &Config{
		DBHost:      getEnvRequired("DB_HOST"),
		DBPort:      getEnvRequired("DB_PORT"),
		DBUser:      getEnvRequired("DB_USER"),
		DBPassword:  getEnvRequired("DB_PASSWORD"),
		DBName:      getEnvRequired("DB_NAME"),
		Port:        getEnvRequired("PORT"),
		FilesAPIURL: getEnvRequired("FILES_API_URL"),
	}
}

func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Required environment variable " + key + " is not set")
	}
	return value
}
