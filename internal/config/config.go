package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	Port      string
	DBConfig  *DBConfig
	ReportURI string
}

type DBConfig struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
	Driver   string `default:"pgx"`
}

const envFile = ".env"

var appConfig *AppConfig = nil
var dbConfig *DBConfig = nil

func init() {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("%s file does not exist...\napp will be stating with default configuration", envFile)
	}
	dbConfig = initDBConfig()
	appConfig = initAppConfig(dbConfig)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func initAppConfig(db *DBConfig) *AppConfig {
	return &AppConfig{
		Port:      getEnv("APP_PORT", "8080"),
		DBConfig:  db,
		ReportURI: "/reports",
	}
}

func initDBConfig() *DBConfig {
	return &DBConfig{
		Port:     getEnv("DB_PORT", "5432"),
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", "test_db"),
		Password: getEnv("DB_PASS", "test_db"),
		DBName:   getEnv("DB_NAME", "test_db"),
		Driver:   getEnv("DB_DRIVER", "pgx"),
	}
}

func GetAppConfig() *AppConfig {
	return appConfig
}

func GetDBConfig() *DBConfig {
	return dbConfig
}
