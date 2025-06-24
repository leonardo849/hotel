package main

import (
	"hotel/internal/logger"
	"hotel/internal/repository"
	"hotel/internal/router"
	"log"
	"os"
)

func main() {
	if _, err :=  repository.ConnectToDatabase();err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	if err := logger.StartLogger(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	if err := router.RunServer(); err != nil {
		logger.ZapLogger.Error(err.Error())
		os.Exit(1)
	}
}

