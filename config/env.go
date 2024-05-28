package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
}

var Envs = initConfig() // global variable to access env variables

func initConfig() Config {
	godotenv.Load() // necessary if you have saved variables in .env file

	// JWT_EXP is in int form in .env file, so need to convert it to string to use it
	jwtExp, _ := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)

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
		PublicHost:             os.Getenv("PUBLIC_HOST"),
		Port:                   os.Getenv("PORT"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBAddress:              os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:                 os.Getenv("DB_NAME"),
		JWTExpirationInSeconds: jwtExp,
	}

}

// necessary if returning default config value
// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}

// 	return fallback
// }

// func getEnvAsInt (key string, fallback int64) int64 {
// 	if value, ok := os.LookupEnv(key); ok {
// 		i, err := strconv.ParseInt(value, 10, 64)
// 		if err != nil {
// 			return fallback
// 		}

// 		return i
// 	}

// 	return fallback
// }
