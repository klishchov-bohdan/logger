package logger

import (
	"log"
	"os"
	"time"
)

type LogToFile struct {
	File *os.File
}

func CreateLogToFile(filePath string) (*LogToFile, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &LogToFile{
		File: file,
	}, nil
}

func (ltf *LogToFile) ClearLogs() error {
	err := ltf.File.Truncate(0)
	if err != nil {
		return err
	}
	return nil
}

//CloseFile closes log file
func (ltf *LogToFile) CloseFile() error {
	err := ltf.File.Close()
	if err != nil {
		return err
	}
	return nil
}

// Info write the info log to the file
func (ltf *LogToFile) Info(msg string) {
	_, err := ltf.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " INFO: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}

// Warning write the warning log to the file
func (ltf *LogToFile) Warning(msg string) {
	_, err := ltf.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " WARNING: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}

// Error write the error log to file
func (ltf *LogToFile) Error(msg string) {
	_, err := ltf.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " ERROR: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}

// Debug write the debug log to file
func (ltf *LogToFile) Debug(msg string) {
	_, err := ltf.File.WriteString(time.Now().Format("01-02-2006 15:04:05") + " DEBUG: " + msg + "\n")
	if err != nil {
		log.Println(err)
	}
}
