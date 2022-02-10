package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	for i := 5; i < 10; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(s1)

	s2 := []int{10, 11, 12, 13}
	/*
		for _, i := range s2 {
			s1 = append(s1, i)
		}
	*/
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
