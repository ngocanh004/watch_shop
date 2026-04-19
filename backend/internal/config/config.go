package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	GinMode    string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSslmode  string

	JWTSecret      string
	JWTExpireHours int

	UploadDir string
	PublicURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Khong tim thay file .env, su dung bien moi truong he thong")
	}

	expire, _ := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "72"))

	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		GinMode:        getEnv("GIN_MODE", "debug"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "watch_shop"),
		DBSslmode:      getEnv("DB_SSLMODE", "disable"),
		JWTSecret:      getEnv("JWT_SECRET", "secret"),
		JWTExpireHours: expire,
		UploadDir:      getEnv("UPLOAD_DIR", "./uploads"),
		PublicURL:      getEnv("PUBLIC_URL", "http://localhost:8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
