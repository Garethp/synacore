package register

var registers[8] int
var stack[] int

func PutRegistry(regIndex, value int) {
	registers[regIndex] = value
}

func GetRegistry(regIndex int) int {
	return registers[regIndex]
}

func PopStack() int {
	var value int
	value, stack = stack[len(stack) - 1], stack[:len(stack) - 1]


	return value
}

func PushStack(value int) {
	stack = append(stack, value)
}