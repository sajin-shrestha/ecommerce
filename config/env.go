package config

import (
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

var Envs = initConfig() // global variable to access env variables

func initConfig() Config {
	godotenv.Load() // necessary if you have saved variables in .env file

	// return Config{
	// 	PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
	// 	Port:       getEnv("PORT", "8080"),
	// 	DBUser:     getEnv("DB_USER", "sajin"),
	// 	DBPassword: getEnv("DB_PASSWORD", "sajinshrestha"),
	// 	DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	// 	DBName:     getEnv("DB_NAME", "ecommerce"),
	// }

	// Initialize Config struct with values from the environment
	return Config{
		PublicHost: os.Getenv("PUBLIC_HOST"),
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

}

// necessary if returning default config value
// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}

// 	return fallback
// }
