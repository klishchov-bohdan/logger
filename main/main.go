package main

import (
	"github.com/klishchov-bohdan/logger"
	"log"
)

func main() {
	fileLogger, err := logger.CreateLogToFile("logs/logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = fileLogger.ClearLogs()
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger := &logger.LogToConsole{}

	logTest(fileLogger)
	logTest(consoleLogger)

	err = fileLogger.CloseFile()
	if err != nil {
		log.Fatal(err)
	}
}

func logTest(lgr logger.Logger) {
	lgr.Info("Created account user@gmail.com")
	lgr.Error("Can`t create user user@gmail.com")
	lgr.Warning("User user@gmail.com can`t login more than 3 times")
	lgr.Debug("User user@gmail.com followed to MAIN page")
}
