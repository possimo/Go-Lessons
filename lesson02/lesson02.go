package main

import (
	"fmt"
	"lesson01/math"
)

var int1, int2 int

func main() {
	int1 = 42
	int2 = 13
	fmt.Printf("%d + %d = %d\n", int1, int2, math.Add(int1, int2))
	fmt.Printf("%d - %d = %d\n", int1, int2, math.Subtract(int1, int2))

	float1, float2 := 13.2, 2.3
	fmt.Println(multiply(float1, float2))
}

func multiply(x, y float64) float64 {
	return x * y
}
