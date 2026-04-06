package config

import (
	"fmt"
	"os"
	"strconv"

	env "github.com/joho/godotenv"
)

func Load() {
	err := env.Load()

	if err != nil {
		fmt.Println("Error in loading dotenv file")
		return
	}
	fmt.Println("Loaded dotenv successfully")
}

func GetString(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Failed to get " + key + "from ENV")
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Failed to get " + key + "from ENV")
		return fallback
	}

	parsedVal, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Failed to parse value to Integer")
		return fallback
	}

	return parsedVal
}

func GetBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Failed to get " + key + "from ENV")
		return fallback
	}

	parseVal, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Println("Failed to parse value to bool")
		return fallback
	}
	return parseVal
}
