package main

import (
	"fmt"
	"../memory"
	"../register"
	//"../log"
	//"os"
	//"bufio"
	//"flag"
	//"log"
	//"flag"
	//"unicode"
	"bytes"
)

var modulo int = 1000000000
//var modulo int = 32768
var commands[] string
var prints bytes.Buffer
var currentOp int

func main() {
	run()

	fmt.Print(commands)
}

func run() {
	if !memory.IsEOM() {
		currentOp = memory.GetMemoryPointer()
		storeOp(getLiteralValue(memory.GetNextMemory()))
		run()
	}
}

func storeOp(opNum int) {
	switch opNum {
	case 0:
		halt()
	case 1:
		setRegistry()
	case 2:
		pushToStack()
	case 3:
		popFromStack()
	case 4:
		equalTo()
	case 5:
		greaterThan()
	case 6:
		jumpTo()
	case 7:
		jumpToIfNotZero()
	case 8:
		jumpToIfZero()
	case 9:
		add()
	case 10:
		multiply()
	case 11:
		mod()
	case 12:
		bitwiseAnd()
	case 13:
		bitwiseOr()
	case 14:
		bitwiseNot()
	case 15:
		readMemory()
	case 16:
		writeToMemory()
	case 17:
		call()
	case 18:
		returnTo()
	case 19:
		printCharacter()
	case 20:
		readCharacter()
	case 21:
		noOp()
	}
}

func halt() {
	addOp("halt")
}


func setRegistry() {
	var registerIndex, value int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory())

	addOp("set", registerIndex % modulo, value)
}


func pushToStack() {
	var a int = getLiteralValue(memory.GetNextMemory())

	addOp("push", a)
}

func popFromStack() {
	var regIndex int = memory.GetNextMemory()

	addOp("pop", regIndex % modulo)
}

func equalTo() {
	var regIndex, a, b = memory.GetNextMemory(), memory.GetNextMemory(), memory.GetNextMemory()
	a = getLiteralValue(a)
	b = getLiteralValue(b)

	addOp("eq", regIndex % modulo, a, b)
}

func greaterThan() {
	var regIndex, a, b = memory.GetNextMemory(), memory.GetNextMemory(), memory.GetNextMemory()
	a = getLiteralValue(a)
	b = getLiteralValue(b)

	addOp("gt", regIndex % modulo, a, b)
}

func jumpTo() {
	var jumpTo int = getLiteralValue(memory.GetNextMemory())

	addOp("jmp", jumpTo)
}

func jumpToIfNotZero() {
	var a, b int = getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("jumpIfNotZero", b, a)
}

func jumpToIfZero() {
	var a int = getLiteralValue(memory.GetNextMemory())
	var b int = memory.GetNextMemory()

	addOp("jumpIfZero", b, a)
}

func add() {
	var reg, a, b = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("add", reg % modulo, a, b)
}

func multiply() {
	var reg, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("mult", reg, a, b)
}

func mod() {
	var reg, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("mod", reg, a, b)
}

func bitwiseAnd() {
	var reg, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("and", reg, a, b)
}

func bitwiseOr() {
	var reg, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("or", reg, a, b)
}

func bitwiseNot() {
	var reg, a int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory())

	addOp("not", reg, a)
}

func readMemory() {
	var registry, memAddress int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory())

	addOp("rmem", registry % modulo, memAddress)
}

func writeToMemory() {
	var memAddress, value = getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	addOp("wmem", memAddress, value)
}

func call() {
	var jump, nextInstruction int = getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetMemoryPointer())

	addOp("call", jump, nextInstruction)
}

func returnTo() {
	addOp("ret")
}

func printCharacter() {
	char := getLiteralValue(memory.GetNextMemory())
	addOp(fmt.Sprintf("out"), char)
}

func readCharacter() {
	var regIndex = memory.GetNextMemory()

	addOp("in", regIndex % modulo)
}

func noOp() {
	addOp("noOp")
}

func addOp(op string, arguments ...int) {
	if op =="out" {
		if (prints.Len() == 0) {
			prints.WriteString(fmt.Sprintf("#%v: out \"", currentOp))
		}
		prints.WriteString(fmt.Sprintf("%c", arguments[0]))
		return
	} else if prints.Len() != 0 {
		commands = append(commands, fmt.Sprintf("%s\"\n", prints.String()))
		prints.Reset()
	}

	commands = append(commands, fmt.Sprintf("#%v: %s %v\n", currentOp, op, arguments))
}

func getLiteralValue(value int) int {
	if (value < modulo) {
		return value
	}

	return value

	register.GetRegistry(0)
	var regIndex int = value % modulo
	return register.GetRegistry(regIndex)
	return regIndex
}

