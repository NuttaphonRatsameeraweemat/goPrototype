package logmanager

import (
	"log"
	"os"
	"time"

	logger "github.com/sirupsen/logrus"
)

func initLog(file *os.File) {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(setFormatter())

	// Set log level.
	logger.SetLevel(logger.InfoLevel)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(file)
}

func getFile(fileName string, pathName string) *os.File {
	path := "Logs/" + pathName + "/"
	os.MkdirAll(path, os.ModePerm)
	file, err := os.OpenFile(path+fileName+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func setFormatter() logger.Formatter {
	Formatter := new(logger.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02T15:04:05.999999999Z07:00"
	Formatter.FullTimestamp = true
	return Formatter
}

// LogInfo func(message string)
func LogInfo(message string) {
	currentTime := time.Now()
	file := getFile(currentTime.Format("2006-01-02"), currentTime.Format("2006-01"))
	initLog(file)
	logger.Info(message)
}

// LogWarning func(message string)
func LogWarning(message string) {
	currentTime := time.Now()
	file := getFile(currentTime.Format("2006-01-02")+"-warning", currentTime.Format("2006-01"))
	initLog(file)
	logger.Warn(message)
}

// LogError func(message string)
func LogError(message string) {
	currentTime := time.Now()
	file := getFile(currentTime.Format("2006-01-02")+"-error", currentTime.Format("2006-01"))
	initLog(file)
	logger.Error(message)
}
