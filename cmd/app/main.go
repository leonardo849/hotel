package main

import (
	"hotel/config"
	"hotel/internal/logger"
	"hotel/internal/repository"
	"hotel/internal/router"
	"hotel/internal/validator"
	"log"
	"os"
)

func main() {

	if err := config.SetupEnvVar(); err != nil {
		log.Fatal(err.Error())
	}

	if err := logger.StartLogger(); err != nil {
		log.Fatal(err.Error())
	}

	validator.StartValidate()

	if _, err :=  repository.ConnectToDatabase();err != nil {
		logger.ZapLogger.Error(err.Error())
		os.Exit(1)
	}
	
	if err := router.RunServer(); err != nil {
		logger.ZapLogger.Error(err.Error())
		os.Exit(1)
	}
	
}

