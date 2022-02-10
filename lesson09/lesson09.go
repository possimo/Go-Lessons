package main

import (
	"fmt"
)

func main() {

	// Pass function to another function
	useFunction(addition, 2, 6)

	// Return function from another function
	calc := getFunction("multiply")
	answer := calc(6, 2)
	fmt.Println("The answer is:", answer)
}

func addition(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func useFunction(calculation func(int, int) int, number1 int, number2 int) {
	answer := calculation(number1, number2)
	fmt.Println("The answer is:", answer)
}

func getFunction(operation string) func(int, int) int {
	var calculation func(int, int) int
	if operation == "add" {
		calculation = func(num1 int, num2 int) int {
			result := num1 + num2
			return result
		}
	} else if operation == "multiply" {
		calculation = func(num1 int, num2 int) int {
			result := num1 * num2
			return result
		}
	}
	return calculation
}
