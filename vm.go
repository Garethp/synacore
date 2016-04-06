package main

import (
	"fmt"
	"./memory"
	"./register"
	"./log"
	"os"
	"bufio"
)

var modulo int = 32768
var halted bool = false
var debug bool = false
var inBuffer[] byte

func main() {
	run()
}

func run() {
	if (shouldKeepRunning()) {
		doOp(getLiteralValue(memory.GetNextMemory()))
		run()
	}
}

func doOp(opNum int) {
	if (debug) {
		//log.Log(fmt.Sprintf("Executing OpCode %v", opNum))
	}

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
	default:
		fmt.Println("Cannot handle OpCode ", opNum)
		halt()
	}
}

func halt() {
	halted = true
}

func setRegistry() {
	var registerIndex, value int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory())

	putValueInRegistry(registerIndex, value)
}

func pushToStack() {
	var a int = getLiteralValue(memory.GetNextMemory())

	register.PushStack(a)
}

func popFromStack() {
	var regIndex, value int = memory.GetNextMemory(), register.PopStack()

	putValueInRegistry(regIndex, value)
}

func equalTo() {
	var regIndex, a, b = memory.GetNextMemory(), memory.GetNextMemory(), memory.GetNextMemory()
	a = getLiteralValue(a)
	b = getLiteralValue(b)

	if (debug) {
		var point int = memory.GetMemoryPointer()
		log.Log(fmt.Sprintf("Comparing %v (%v) and %v (%v) putting them in to registry %v\r\n", a, point - 2, b, point - 1, regIndex % modulo))
	}

	if (a == b) {
		putValueInRegistry(regIndex, 1)
	} else {
		putValueInRegistry(regIndex, 0)
	}
}

func greaterThan() {
	var regIndex, a, b = memory.GetNextMemory(), memory.GetNextMemory(), memory.GetNextMemory()
	a = getLiteralValue(a)
	b = getLiteralValue(b)

	if (a > b) {
		putValueInRegistry(regIndex, 1)
	} else {
		putValueInRegistry(regIndex, 0)
	}
}

func jumpTo() {
	var jumpTo int = getLiteralValue(memory.GetNextMemory())

	if (debug) {
		log.Log(fmt.Sprintf("Jumping to %v (%v)", jumpTo, memory.GetMemoryPointer() - 1))
	}

	memory.SetMemoryPointer(jumpTo)
}

func jumpToIfNotZero() {
	var a, b int = getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	if (getLiteralValue(a) != 0) {
		memory.SetMemoryPointer(b)
	}
}

func jumpToIfZero() {
	var a int = getLiteralValue(memory.GetNextMemory())
	var b int = memory.GetNextMemory()

	if (getLiteralValue(a) == 0){
		memory.SetMemoryPointer(b)
	}
}

func add() {
	var reg, a, b = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())
	var result = (a + b) % modulo

	if (reg % modulo == 0 && result == 411 && debug) {
		var point int = memory.GetMemoryPointer()
		log.Log(fmt.Sprintf("Adding %v (%v) and %v (%v) = %v putting them in to registry %v\r\n", a, point - 2, b, point - 1, result, reg % modulo))
	}

	putValueInRegistry(reg, result)
}

func multiply() {
	var register, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())
	var result int = (a * b) % modulo

	if (debug) {
		var point int = memory.GetMemoryPointer()
		log.Log(fmt.Sprintf("Multiplying %v (%v) with %v (%v) to get %v and putting it in to Reg %v\r\n", a, point - 2, b, point - 1, result, register % modulo))
	}

	putValueInRegistry(register, result)
}

func mod() {
	var register, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	putValueInRegistry(register, a % b)
}

func bitwiseAnd() {
	var register, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	a16, b16 := uint16(a), uint16(b)

	putValueInRegistry(register, int(a16 & b16) % modulo)
}

func bitwiseOr() {
	var register, a, b int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	a16, b16 := uint16(a), uint16(b)

	putValueInRegistry(register, int(a16 | b16) % modulo)
}

func bitwiseNot() {
	var register, a int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory())

	var a16 = uint16(a)

	putValueInRegistry(register, int(^a16) % modulo)
}

func readMemory() {
	var registry, memAddress int = memory.GetNextMemory(), getLiteralValue(memory.GetNextMemory())

	putValueInRegistry(registry, memory.Read(memAddress))
}

func writeToMemory() {
	var memAddress, value = getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetNextMemory())

	memory.Write(memAddress, value)
}

func call() {
	var jump, nextInstruction int = getLiteralValue(memory.GetNextMemory()), getLiteralValue(memory.GetMemoryPointer())

	register.PushStack(nextInstruction)
	memory.SetMemoryPointer(getLiteralValue(jump))
}

func returnTo() {
	var jump int = getLiteralValue(register.PopStack())

	memory.SetMemoryPointer(jump)
}

func printCharacter() {
	fmt.Printf("%c", getLiteralValue(memory.GetNextMemory()))
}

func readCharacter() {
	var regIndex = memory.GetNextMemory()

	if (len(inBuffer) == 0) {
		inBuffer = readLine()
	}

	var code int
	code, inBuffer = int(inBuffer[0]), inBuffer[1:]

	putValueInRegistry(regIndex, int(code))
}

func noOp() {

}

func autoPlay() []byte {
	var commands[] string = []string{
		"take tablet",
		"doorway",
		"north",
		"north",
		"bridge",
		"continue",
		"down",
		"east",
		"take empty lantern",
		"west",
		"west",
		"passage",
		"ladder",
		"west",
		"south",
		"north",
		"take can",
		"west",
		"ladder",
		"darkness",
		"use can",
		"use lantern",
		"continue",
		"west",
		"west",
		"west",
		"west",
		"north",
		"take red coin",
		"north",
		"east",
		"take concave coin",
		"down",
		"take corroded coin",
		"up",
		"west",
		"west",
		"take blue coin",
		"up",
		"take shiny coin",
		"down",
		"east",
		"use blue coin",
		"use red coin",
		"use shiny coin",
		"use concave coin",
		"use corroded coin",
		"north",
		"take teleporter",
		"use teleporter",
		"take business card",
		"take strange book",
		""}

	var autoplay[] byte

	for index, command := range commands {
		if (command == "") {
			continue
		}

		index = index
		command = command + "\n"

		var commandBytes[] byte = []byte(command)
		autoplay = append(autoplay, commandBytes[:]...)
	}
	return autoplay
}

func readLine() []byte {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var line = []byte(text)
	line = stripWindowsCarriageReturn(line)

	if (string(line) == "start debug\n") {
		debug = true
		fmt.Println("Debug on")
		return readLine()
	} else if (string(line) == "stop debug\n") {
		debug = false
		fmt.Println("Debug on")
		return readLine()
	} else if (string(line) == "dump registry\n") {
		fmt.Println(register.GetRegisters())
		return readLine()
	} else if (string(line) == "dump stack\n") {
		fmt.Println(register.GetStack())
		return readLine()
	} else if (string(line) == "autoplay\n") {
		return autoPlay()
	}

	return line
}

func getLiteralValue(value int) int {
	if (value < modulo) {
		return value
	}

	var regIndex int = value % modulo
	return register.GetRegistry(regIndex)
}

func putValueInRegistry(fullReg, value int) {
	var regIndex int = fullReg % modulo

	register.PutRegistry(regIndex, value)
}

func shouldKeepRunning() bool {
	return !halted && !memory.IsEOM()
}

func stripWindowsCarriageReturn(line[] byte) []byte {
	//Windows terminals send a [13 10] for \r\n instead of just [10] for \n. We're looking for this and stripping
	if (string(line[len(line) - 2:]) == "\r\n") {
		line = append(line[:len(line) - 2], line[len(line) - 1:]...)
	}

	return line
}