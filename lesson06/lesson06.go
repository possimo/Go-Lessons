package main

import "fmt"

type person struct {
	name    string
	age     int
	married bool
	contact
}

type contact struct {
	mobile string
	office string
	email  string
}

func main() {
	personSlice := createPersonSlice()
	personSlice[0].contact.email = "ali@possimo.com"
	personSlice[1].contact.mobile = "0123456789"
	personSlice[2].contact.office = "0334567890"
	personSlice[2].email = ""

	for i, p := range personSlice {
		fmt.Println("index:", i, "person:", p)
	}

	newPerson1 := person{name: "Ina", age: 15}
	personSlice = append(personSlice, newPerson1)
	fmt.Println("personSlice Length:", len(personSlice), " Capacity:", cap(personSlice))

	for _, p := range personSlice {
		fmt.Println(p.name, "is", p.age, "years old")
	}

	newPerson2 := person{name: "Eli", age: 5}
	addToPersonSlice(newPerson2, &personSlice)

	for i := range personSlice {
		//fmt.Println(personSlice[i:])
		//fmt.Println(personSlice[:i])
		fmt.Println(personSlice[i : i+1])
	}

	deleteFromPersonSlice("Abu", &personSlice)
	fmt.Println(personSlice)
}

func createPersonSlice() []person {
	//var personSlice []person
	/*
		personSlice := make([]person, 2, 3)

		fmt.Println("personSlice Length:", len(personSlice), " Capacity:", cap(personSlice))
		//personSlice = append(personSlice, person{name: "Ali", age: 20})
		personSlice[0] = person{name: "Ali", age: 20}

		fmt.Println("personSlice Length:", len(personSlice), " Capacity:", cap(personSlice))
		//personSlice = append(personSlice, person{name: "Umi", age: 40, married: true})
		personSlice[1] = person{name: "Umi", age: 40, married: true}

		fmt.Println("personSlice Length:", len(personSlice), " Capacity:", cap(personSlice))
		personSlice = append(personSlice, person{name: "Abi", age: 50, married: true})
	*/

	personSlice := []person{
		{name: "Ali", age: 20},
		{name: "Umi", age: 40, married: true},
		{name: "Abi", age: 50, married: true},
	}

	fmt.Println("personSlice Length:", len(personSlice), " Capacity:", cap(personSlice))
	return personSlice
}

func addToPersonSlice(newPerson person, personSlice *[]person) {
	*personSlice = append(*personSlice, newPerson)
	fmt.Println("personSlice Length:", len(*personSlice), " Capacity:", cap(*personSlice))
}

func deleteFromPersonSlice(name string, personSlice *[]person) {
	for i, p := range *personSlice {
		if p.name == name {
			*personSlice = append((*personSlice)[:i], (*personSlice)[i+1:]...)
			return
		}
	}
	fmt.Println("Person:", name, "not found in PersonSlice")
}

/*
Note:
“In Go, a slice grows by doubling until it contains 1024 elements,
after which it grows by 25% each time”

Excerpt From
100 Go Mistakes and How to Avoid Them MEAP V09
Teiva Harsanyi
This material may be protected by copyright.
*/
