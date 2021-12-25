package logger

import (
	"log"
	"time"
)

type LogToConsole struct {
}

// Info write the info log to the console
func (ltc *LogToConsole) Info(msg string) {
	log.Println(time.Now().Format("01-02-2006 15:04:05") + " INFO: " + msg)
}

// Warning write the warning log to the console
func (ltc *LogToConsole) Warning(msg string) {
	log.Println(time.Now().Format("01-02-2006 15:04:05") + " WARNING: " + msg)
}

// Error write the error log to console
func (ltc *LogToConsole) Error(msg string) {
	log.Println(time.Now().Format("01-02-2006 15:04:05") + " ERROR: " + msg)
}

// Debug write the debug log to console
func (ltc *LogToConsole) Debug(msg string) {
	log.Println(time.Now().Format("01-02-2006 15:04:05") + " DEBUG: " + msg)
}
