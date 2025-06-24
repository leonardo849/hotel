package main

import (
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
	if err := router.RunServer(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	router.RunServer()
}

