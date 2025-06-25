package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SetupEnvVar() error {
	mode := os.Getenv("APP_ENV")
	if mode == "" || mode == "DEV" {
		cwd, _ := os.Getwd()
		projectRoot := filepath.Join(cwd, "..", "..")
		envPath := filepath.Join(projectRoot, "config", ".env")
		err := godotenv.Load(envPath)
		log.Print("dev mode")
		if err != nil {
			return err
		}
	} else {
		log.Print("prod mode")
	}
	
	return  nil
}