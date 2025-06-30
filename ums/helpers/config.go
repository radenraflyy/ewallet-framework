package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetUpConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func GetEnv(key, defVal string) string {
	result, exists := Env[key]
	if !exists {
		result = defVal
	}
	return result
}
