package main

import "fmt"

type person struct {
	name    string
	age     int
	married bool
}

type contact struct {
	code   string
	email  string
	mobile string
}

func readIface(i interface{}) {
	// show param value
	fmt.Printf("%#v %T\n", i, i)

	// use type assertion to check for string  type
	if s, ok := i.(string); ok {
		fmt.Println("string:", s)
	}

	//  use type assertion to check for person type
	if p, ok := i.(person); ok {
		fmt.Println("person:", p)
	}

	//  use type assertion to check for contact type
	if c, ok := i.(contact); ok {
		fmt.Println("contact:", c)
	}
}

func readXface(i interface{}) {
	// show param value
	fmt.Printf("%#v %T\n", i, i)

	switch v := i.(type) {
	case string:
		fmt.Println("string:", v)
	case person:
		fmt.Println("person:", v)
	case contact:
		fmt.Println("contact:", v)
	}
}

func main() {

	msg := "hello world"
	//readIface(msg)
	readXface(msg)

	ali := person{name: "Ali", age: 20, married: true}
	//readIface(msg)
	readXface(ali)

	spy := contact{code: "001", email: "ejenali@possimo.com", mobile: "0123456789"}
	//readIface(msg)
	readXface(spy)
}

/*
Notes:

Every type will implement the empty interface: interface{},
since every type will have at least no method.
So every type can qualify as the empty interface — interface{} type.

- excerpt from "Understanding Interfaces in Go" by Arun Kant Pant
*/
