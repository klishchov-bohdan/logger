package main

import (
	"github.com/klishchov-bohdan/logger"
	"log"
)

func main() {
	fileLogger, err := logger.NewFileLogger("logs/logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(fileLogger *logger.FileLogger) {
		err := fileLogger.CloseFile()
		if err != nil {
			log.Fatal(err)
		}
	}(fileLogger)
	err = fileLogger.ClearLogs()
	if err != nil {
		log.Fatal(err)
	}

	consoleLogger := logger.NewConsoleLogger()

	logTest(fileLogger)
	logTest(consoleLogger)
}

func logTest(lgr logger.Logger) {
	lgr.Info("Created account user@gmail.com")
	lgr.Error("Can`t create user user@gmail.com")
	lgr.Warning("User user@gmail.com can`t login more than 3 times")
	lgr.Debug("User user@gmail.com followed to MAIN page")
}
