package main

import "fmt"

type getInfoType func(string, int, bool) string

type person struct {
	name    string
	age     int
	married bool
	getInfo getInfoType
}

func main() {
	ali := person{
		name:    "Ali",
		age:     20,
		married: true,
		getInfo: func(name string, age int, married bool) string {
			return fmt.Sprintf("%s is %d years old and married status is %t", name, age, married)
		},
	}
	fmt.Println(ali.getInfo(ali.name, ali.age, ali.married))
}
