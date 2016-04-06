package log

import (
	"os"
	//"fmt"
)

var opened bool = false
var logFile *os.File

func getLogFile() *os.File {
	if (!opened) {
		os.Remove("./log.log")
		file, error := os.Create("./log.log")
		file.Close()

		logFile, error = os.OpenFile("./log.log", os.O_APPEND|os.O_WRONLY, 0600)

		if (error != nil) {
			panic(error)
		}

		opened = true
	}

	return logFile
}

func Log(message string) {
	file := getLogFile()

	message = message + "\r\n"

	if _, err := file.WriteString(message); err != nil {
		panic(err)
	}
}