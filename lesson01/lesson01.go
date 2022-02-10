package main

import (
	"fmt"
	"lesson01/math"
)

func main() {
	fmt.Println("42 + 13 =", add(42, 13))
	fmt.Println("42 + 13 =", math.Add(42, 13))
	fmt.Println("42 - 13 =", math.Subtract(42, 13))
}

func add(x int, y int) int {
	return x + y
}
