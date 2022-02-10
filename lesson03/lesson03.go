package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	s := "Possimo Technologies"
	printChar(s)
}

func printChar(s string) {
	for _, char := range s {
		fmt.Println(char, "\t->", string(char))
	}
}
