package config

import (
	"os"
	"strconv"
)

type DataBaseConf struct {
	Host         string
	Port         int
	DataBaseName string
	UserName     string
	Password     string
}

type Config struct {
	DataBase DataBaseConf
}

func NewConfig() *Config {
	return &Config{
		DataBase: DataBaseConf{
			Host:         getEnv("DATABASE_HOST", "localhost"),
			Port:         getEnvInt("DATABASE_PORT", 5432),
			DataBaseName: getEnv("POSTGRES_DB", "postgres"),
			UserName:     getEnv("POSTGRES_USER", "root"),
			Password:     getEnv("POSTGRES_PASSWORD", "password"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		result, err := strconv.Atoi(value)
		if err != nil {
			return defaultVal
		}
		return result
	}

	return defaultVal
}
