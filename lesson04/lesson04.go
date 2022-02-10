package main

import "fmt"

type person struct {
	name    string
	age     int
	married bool
}

func main() {
	/*
		var ati person
		ati.name = "ati"
		ati.age = 20
		ati.married = true
		fmt.Println(ati)
	*/
	//ali := person{"Ali", 30, false}
	ali := person{name: "Ali", age: 30, married: true}
	/*
		fmt.Println(ali)
		ali.married = false
		fmt.Println(ali)

		if ali.married {
			fmt.Println(ali.name, ali.age, "years old is married")
		} else {
			fmt.Println(ali.name, ali.age, "years old is not married")
		}
	*/

	p := &ali
	fmt.Println(*p)

	p.age = 40
	fmt.Println(ali)

	fmt.Println(printPerson(ali))
}

func printPerson(p person) string {
	var details struct {
		name string
		age  string
	}

	details.name = p.name
	details.age = fmt.Sprintf("%d", p.age)

	s := "My name is " + details.name + " and I am " + details.age + " years old"
	return s
}
