package config

import (
	"cmp"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort   string
	MongoURI  string
	DBName    string
	LogLevel  string
	LogFormat string
	DebugMode bool
}

var (
	cfg  *AppConfig
	once sync.Once
)

func LoadConfig() {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Print("Error loading .env file", err)
		}

		cfg = &AppConfig{
			AppPort:   getEnv("APP_PORT", "8080"),
			MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
			DBName:    getEnv("DB_NAME", "bookstore"),
			LogLevel:  getEnv("LOG_LEVEL", "debug"),
			LogFormat: getEnv("LOG_FORMAT", "console"),
		}
		cfg.DebugMode = cfg.LogLevel == "debug"
		log.Printf("Config loaded: %+v\n", *cfg)
	})
}

func Get() *AppConfig {
	if cfg == nil {
		LoadConfig()
	}
	return cfg
}

// Helper to fetch env vars with default
func getEnv(key, fallback string) string {
	return cmp.Or(os.Getenv(key), fallback)
}
