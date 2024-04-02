package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_User", "pimo"),
		DBPassword: getEnv("DB_Password", "9689x"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_Host", "127.0.0.1"), getEnv("DB_Port", "3306")),
		DBName:     getEnv("DB_Name", "form"),
	}
}
