package config

import (
	"log"
	"os"
	"github.com/joho/godotenv" // Requiere: go get github.com/joho/godotenv
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ℹ️ No se encontró archivo .env, usando variables de entorno del sistema.")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}