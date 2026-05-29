package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	AppPort           string
	AppEnv            string
	JwtAccessSecret   string
	JwtRefreshSecret  string
	JwtAccessExpMin   int
	JwtRefreshExpDays int
	YandexClientID    string
	YandexClientSecret string
	YandexRedirectURI string
}

var AppConfig *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, reading from environment")
	}

	accessExpMin, _ := strconv.Atoi(getEnv("JWT_ACCESS_EXPIRATION_MIN", "15"))
	refreshExpDays, _ := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRATION_DAYS", "7"))

	AppConfig = &Config{
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "student"),
		DBPassword:         getEnv("DB_PASSWORD", "student_secure_password"),
		DBName:             getEnv("DB_NAME", "wp_labs"),
		AppPort:            getEnv("PORT", "4200"),
		AppEnv:             getEnv("APP_ENV", "development"),
		JwtAccessSecret:    getEnv("JWT_ACCESS_SECRET", "change_me_access"),
		JwtRefreshSecret:   getEnv("JWT_REFRESH_SECRET", "change_me_refresh"),
		JwtAccessExpMin:    accessExpMin,
		JwtRefreshExpDays:  refreshExpDays,
		YandexClientID:     getEnv("YANDEX_CLIENT_ID", ""),
		YandexClientSecret: getEnv("YANDEX_CLIENT_SECRET", ""),
		YandexRedirectURI:  getEnv("YANDEX_REDIRECT_URI", "http://localhost:4200/auth/oauth/yandex/callback"),
	}
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}