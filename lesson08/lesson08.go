package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Use function as variable
	var add func(int, int) int
	fmt.Println(reflect.TypeOf(add))
	add = addition

	//add := addition
	//fmt.Println(reflect.TypeOf(add))

	formula := addition
	fmt.Println(reflect.TypeOf(formula))
	answer := formula(2, 6)
	fmt.Println("2 + 6 =", answer)

	formula = multiplication
	answer = formula(2, 6)
	fmt.Println("2 x 6 =", answer)

	// Create anonymous function as a variable
	greetings := func(username string) {
		fmt.Println("Apa khabar", username)
	}
	fmt.Println(reflect.TypeOf(greetings))

	greetings("Ali")
	greetings("Umi")

}

func addition(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func multiplication(num1 int, num2 int) int {
	result := num1 * num2
	return result
}
