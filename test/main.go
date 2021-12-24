package main

import (
	"github.com/klishchov-bohdan/logger"
	"log"
)

func main() {
	logger, err := logger.CreateLogToFile("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = logger.ClearLogs()
	if err != nil {
		log.Fatal(err)
	}
	err = logger.Info("log")
	if err != nil {
		log.Fatal(err)
	}
	_ = logger.CloseFile()
}
