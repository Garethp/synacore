package memory

import (
	"os"
	"io"
	"bytes"
	"encoding/binary"
	"fmt"
)

var register[] int
var memoryPointer int = 0

func LoadFile() {
	fmt.Println("Reading File")
	var fp, _ = os.Open("./challenge.bin")
	var buf = make([]byte, 2)

	for true {
		_, err := fp.Read(buf)

		if err == io.EOF {
			break
		}

		buffer := bytes.NewBuffer(buf)
		var k uint16
		binary.Read(buffer, binary.LittleEndian, &k)
		register = append(register, int(k))
	}

	fmt.Println("Done")
}

func GetNextMemory() int {
	if (len(register) == 0) {
		LoadFile()
	}

	var value int = register[memoryPointer]
	memoryPointer++
	return value
}

func SetMemoryPointer(pointer int) {
	memoryPointer = pointer
}

func GetMemoryPointer() int {
	return memoryPointer
}

func IsEOM() bool {
	if (len(register) == 0) {
		LoadFile()
	}

	if (memoryPointer <= len(register) - 1) {
		return false
	}

	fmt.Println("Out of Memory", memoryPointer, len(register))
	return true
}

func Write(pointer, value int) {
	register[pointer] = value
}

func Read(pointer int) int {
	return register[pointer]
}