package main

import (
	"fmt"
	"./memory"
	"./register"
	"os"
	"bufio"
)

var modulo int = 32768
var halted bool = false
var debug bool = false
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
	if (debug) {
		fmt.Println(opNum)
	}

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

	if (debug) {
		println("Comparing", a, b)
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

	putValue(reg, result)
}

func multiply() {
	var register, a, b int = memory.GetNextMemory(), getValue(memory.GetNextMemory()), getValue(memory.GetNextMemory())

	putValue(register, (a * b) % modulo)
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
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		readbuffer = []byte(text)
	}

	var code int
	code, readbuffer = int(readbuffer[0]), readbuffer[1:]
	putValue(regIndex, int(code))

	//var b []byte = make([]byte, 2)
	//os.Stdin.Read(b)
	//var code uint16
	//_ = binary.Read(bytes.NewReader(b), binary.LittleEndian, &code)
	//
	//putValue(regIndex, int(code))
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