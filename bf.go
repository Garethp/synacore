package main

import (
	"fmt"
	"math"
)

var numbers[] int = []int{2, 3, 5, 7, 9}
var pointers[] int = []int{0, 1, 2, 3, 4}

var desired int = 399

func main() {
	brute()

	fmt.Println(pointers)
}

func brute() {
	var result int = calculate()
	if (result != desired || uniqueCount(pointers) != 5) {
		incPointer(4)
		brute()
	}
}

func incPointer(index int) {
	pointers[index]++
	if (pointers[index] > 4 && index > 0) {
		incPointer(index - 1)
		pointers[index] = 0
	}
}

func getNumbers() (int, int, int, int, int) {
	return numbers[pointers[0]], numbers[pointers[1]], numbers[pointers[2]], numbers[pointers[3]], numbers[pointers[4]]
}

func calculate() int {
	var a, b, c, d, e int = getNumbers()

	return a + b * int(math.Pow(float64(c), 2)) + int(math.Pow(float64(d), 3)) - e
}

func uniqueCount (toCheck []int) int {
	var foundSoFar[] int

	for index, value := range toCheck {
		index = index
		if (!inArray(value, foundSoFar)) {
			foundSoFar = append(foundSoFar, value)
		}
	}

	return len(foundSoFar)
}

func inArray(needle int, haystack []int) bool {
	for index, value := range haystack {
		index = index
		if (value == needle) {
			return true
		}
	}

	return false
}