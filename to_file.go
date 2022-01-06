package logger

import (
	"log"
	"os"
	"time"
)

type FileLogger struct {
	File *os.File
}

func NewFileLogger(filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{
		File: file,
	}, nil
}

func (fl *FileLogger) ClearLogs() error {
	err := fl.File.Truncate(0)
	if err != nil {
		return err
	}
	return nil
}

//CloseFile closes log file
func (fl *FileLogger) CloseFile() error {
	err := fl.File.Close()
	if err != nil {
		return err
	}
	return nil
}

// Info write the info log to the file
func (fl *FileLogger) Info(msg string) {
	_, err := fl.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " INFO: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}

// Warning write the warning log to the file
func (fl *FileLogger) Warning(msg string) {
	_, err := fl.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " WARNING: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}

// Error write the error log to file
func (fl *FileLogger) Error(msg string) {
	_, err := fl.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " ERROR: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}

// Debug write the debug log to file
func (fl *FileLogger) Debug(msg string) {
	_, err := fl.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " DEBUG: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}
