package log

import (
	"os"
	//"fmt"
	"fmt"
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

func getOpName(opCode int) string {
	operations := []string{"halt", "setRegistry", "pushToStack", "pop", "equalsTo", "greaterThan", "jumpTo", "jumpIfNot0", "jumpIf0", "add", "mul", "mod", "and", "or", "not", "readMem", "writeMem", "call", "returnTo", "printChar", "readChar", "noOp"}

	return operations[opCode]
}

func Log(message string) {
	file := getLogFile()

	message = message + "\r\n"

	if _, err := file.WriteString(message); err != nil {
		panic(err)
	}
}

func LogOpCode(point int, opCode int) {
	file := getLogFile()
	message := fmt.Sprintf("#%v: Executing OpCode %v\r\n", point, getOpName(opCode))

	if _, err := file.WriteString(message); err != nil {
		panic(err)
	}
}