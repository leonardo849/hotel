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
		projectRoot := findProjectRoot()
		if projectRoot == "" {
			return ErrProjectRootNotFound
		}

		envPath := filepath.Join(projectRoot, "config", ".env")
		err := godotenv.Load(envPath)
		log.Print("dev mode")
		if err != nil {
			return err
		}
	} else {
		log.Print("prod mode")
	}
	return nil
}


func findProjectRoot() string {
	dir, _ := os.Getwd() //get current directory
	for { //if it doesn't find the file in dir, it goes up to parent folder
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil { //check if the file exists and get it. 
			return dir
		}
		parent := filepath.Dir(dir) //get the parent
		if parent == dir {
			break 
		}
		dir = parent
	}
	return ""
}


var ErrProjectRootNotFound = os.ErrNotExist
