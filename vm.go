package main

import (
	"fmt"
	"./memory"
	"./register"
	"os"
	"bufio"
	"math"
)

var modulo int = 32768
var halted bool = false
var debugEnabled bool = false
var readbuffer[] byte

func main() {
	run()
}

func run() {
	if (continueRunning()) {
		doOp(getValue(memory.GetNextMemory()))
		run()
	}
}

func doOp(opNum int) {
	switch opNum {
	case 0:
		halt()
	case 1:
		set()
	case 2:
		push()
	case 3:
		pop()
	case 4:
		equalTo()
	case 5:
		greaterThan()
	case 6:
		jump()
	case 7:
		jt()
	case 8:
		jf()
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
		writeMemory()
	case 17:
		call()
	case 18:
		ret()
	case 19:
		printAsAscii()
	case 20:
		readchar()
	case 21:
		noop()
	default:
		fmt.Println("Cannot handle OpCode ", opNum)
		halt()
	}
}

func halt() {
	halted = true
}

func set() {
	var registerIndex, value int = memory.GetNextMemory(), getValue(memory.GetNextMemory())

	putValue(registerIndex, value)
}

func push() {
	var a int = getValue(memory.GetNextMemory())

	register.PushStack(a)
}

func pop() {
	var regIndex, value int = memory.GetNextMemory(), register.PopStack()

	putValue(regIndex, value)
}

func equalTo() {
	var regIndex, a, b = memory.GetNextMemory(), memory.GetNextMemory(), memory.GetNextMemory()
	a = getValue(a)
	b = getValue(b)

	if (debugEnabled && b == 399) {
		var point int = memory.GetMemoryPointer()
		fmt.Println(memory.Read(5235))
		fmt.Println(register.GetRegisters())
		fmt.Printf("Comparing %v (%v) and %v (%v) putting them in to registry %v\r\n", a, point - 2, b, point - 1, regIndex % modulo)
	}

	if (a == b) {
		putValue(regIndex, 1)
	} else {
		putValue(regIndex, 0)
	}
}

func greaterThan() {
	var regIndex, a, b = memory.GetNextMemory(), memory.GetNextMemory(), memory.GetNextMemory()
	a = getValue(a)
	b = getValue(b)

	if (a > b) {
		putValue(regIndex, 1)
	} else {
		putValue(regIndex, 0)
	}
}

func jump() {
	var jumpTo int = getValue(memory.GetNextMemory())
	memory.SetMemoryPointer(jumpTo)
}

func jt() {
	var a, b int = getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())

	if (getValue(a) != 0) {
		memory.SetMemoryPointer(b)
	}
}

func jf() {
	var a int = getValue(memory.GetNextMemory())
	var b int = memory.GetNextMemory()

	if (getValue(a) == 0){
		memory.SetMemoryPointer(b)
	}
}

func add() {
	var reg, a, b = memory.GetNextMemory(), getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())
	var result = (a + b) % modulo

	if (reg % modulo == 0 && result == 411 && debugEnabled) {
		var point int = memory.GetMemoryPointer()
		fmt.Printf("Adding %v (%v) and %v (%v) = %v putting them in to registry %v\r\n", a, point - 2, b, point - 1, result, reg % modulo)
	}

	putValue(reg, result)
}

func multiply() {
	var register, a, b int = memory.GetNextMemory(), getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())
	var result int = (a * b) % modulo

	if (debugEnabled) {
		fmt.Println("Multiplying", a, b)
//		var point int = memory.GetMemoryPointer()
//		fmt.Printf("Multiplyiny %v (%v) with %v (%v) to get %v and putting it in to Reg %v\r\n", a, point - 2, b, point - 1, result, register % modulo)
	}

	putValue(register, result)
}

func mod() {
	var register, a, b int = memory.GetNextMemory(), getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())

	putValue(register, a % b)
}

func bitwiseAnd() {
	var register, a, b int = memory.GetNextMemory(), getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())

	a16, b16 := uint16(a), uint16(b)

	putValue(register, int(a16 & b16) % modulo)
}

func bitwiseOr() {
	var register, a, b int = memory.GetNextMemory(), getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())

	a16, b16 := uint16(a), uint16(b)

	putValue(register, int(a16 | b16) % modulo)
}

func bitwiseNot() {
	var register, a int = memory.GetNextMemory(), getValue(memory.GetNextMemory())

	var a16 = uint16(a)

	putValue(register, int(^a16) % modulo)
}

func readMemory() {
	var registry, memAddress int = memory.GetNextMemory(), getValue(memory.GetNextMemory())

	putValue(registry, memory.Read(memAddress))
}

func writeMemory() {
	var memAddress, value = getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())

	memory.Write(memAddress, value)
}

//If there are problems, look here
func call() {
	var jump, nextInstruction int = getValue(memory.GetNextMemory()), getValue(memory.GetMemoryPointer())

	register.PushStack(nextInstruction)
	memory.SetMemoryPointer(getValue(jump))
}

func ret() {
	var jump int = getValue(register.PopStack())

	memory.SetMemoryPointer(jump)
}

func printAsAscii() {
	fmt.Printf("%c", getValue(memory.GetNextMemory()))
}

func readchar() {
	var regIndex = memory.GetNextMemory()

	if (len(readbuffer) == 0) {
		readbuffer = readLine()
	}

	var code int
	code, readbuffer = int(readbuffer[0]), readbuffer[1:]

	putValue(regIndex, int(code))
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
		"use shiny coin",
		"use corroded coin",
		"use red coin",
		"use concave coin",
		"use blue coin",
		""}

	var autoplay[] byte

	for index, command := range commands {
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
	line = stripWindowsCr(line)

	if (string(line) == "start debug\n") {
		debugEnabled = true
		fmt.Println("Debug on")
		return readLine()
	} else if (string(line) == "stop debug\n") {
		debugEnabled = false
		fmt.Println("Debug on")
		return readLine()
	} else if (string(line) == "dump registry\n") {
		fmt.Println(register.GetRegisters())
		return readLine()
	} else if (string(line) == "dump stack\n") {
		fmt.Println(register.GetStack())
		return readLine()
	} else if (string(line) == "test calc\n") {
//		fmt.Println("Answer", 3 + 5 * math.Pow(9, 2) + math.Pow(2, 3) - 7) 409
		fmt.Println("Answer", 3 + 5 * math.Pow(7, 2) + math.Pow(2, 3) - 9)
		return readLine()
	} else if (string(line) == "autoplay\n") {
		return autoPlay()
	}

	return line
}

func noop() {

}

func getValue(value int) int {
	if (value < modulo) {
		return value
	}

	var regIndex int = value % modulo
	return register.GetRegistry(regIndex)
}

func putValue(fullReg, value int) {
	var regIndex int = fullReg % modulo

	register.PutRegistry(regIndex, value)
}

func continueRunning() bool {
	return !halted && !memory.IsEOM()
}

func stripWindowsCr(readbuffer[] byte) []byte {
	//Windows terminals send a [13 10] for \r\n instead of just [10] for \n. We're looking for this and stripping
	if (int(readbuffer[len(readbuffer) - 2]) == 13 && int(readbuffer[len(readbuffer) - 1]) == 10) {
		var nl byte = readbuffer[len(readbuffer) - 1]
		readbuffer = readbuffer[:len(readbuffer) - 2]
		readbuffer = append(readbuffer, nl)
	}

	return readbuffer
}