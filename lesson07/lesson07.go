package main

import "fmt"

type person struct {
	name    string
	age     int
	married bool
}

var m map[int]string     // nil map
var n = map[int]string{} // empty map

func main() {
	m = make(map[int]string)
	m[1] = "a"
	fmt.Println(m)

	n[1] = "A"
	fmt.Println(n)

	personMap := createPersonMap()
	fmt.Println(personMap)

	personAli := personMap["01"]
	fmt.Println(personAli)

	personMap["04"] = person{name: "Ina", age: 15}
	fmt.Println(personMap)

	delete(personMap, "02")
	fmt.Println(personMap)

	if p, ok := personMap["03"]; ok {
		fmt.Println("Person:", p, "is in personMap")
	}
}

func createPersonMap() map[string]person {
	personMap := map[string]person{
		"01": {name: "Ali", age: 20},
		"02": {name: "Umi", age: 40, married: true},
		"03": {name: "Abi", age: 50, married: true},
	}
	return personMap
}
