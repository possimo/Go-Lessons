package main

import "fmt"

type receiver interface {
	sendSMS(string) error
	sendEmail(string) error
}

type database interface {
	save() error
	delete() error
}

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

type employee struct {
	department string
	person
	contact
}

func (e employee) sendSMS(message string) error {
	fmt.Printf("Sending [%s] to employee:%s via SMS...\n", message, e.mobile)
	return nil
}

func (e employee) sendEmail(message string) error {
	fmt.Printf("Sending [%s] to employee:%s via email...\n", message, e.email)
	return nil
}

func (e employee) save() error {
	fmt.Println("Saving employee:", e.name, "in database...")
	return nil
}

func (e employee) delete() error {
	fmt.Println("Deleting employee:", e.name, "in database...")
	return nil
}

type client struct {
	company string
	person
	contact
}

func (c client) sendSMS(message string) error {
	fmt.Printf("Sending [%s] to client:%s via SMS...\n", message, c.mobile)
	return nil
}

func (c client) sendEmail(message string) error {
	fmt.Printf("Sending [%s] to client:%s via email...\n", message, c.email)
	return nil
}

func sendMessage(receiver receiver, message string) error {
	receiver.sendSMS(message)
	receiver.sendEmail(message)
	return nil
}

func main() {
	var db database

	newEmployee := employee{
		department: "Sales",
		person:     person{name: "Ali", age: 30, married: true},
		contact:    contact{code: "001", email: "ali@possimo.com", mobile: "0123456789"},
	}
	sendMessage(newEmployee, "hello")

	newClient := client{
		company: "Altel",
		person:  person{name: "Umi", age: 40, married: true},
		contact: contact{code: "999", email: "umi@altel.com.my", mobile: "0198765432"},
	}
	sendMessage(newClient, "assalamualaikum")

	db = employee{person: person{name: "Abu", age: 30}}
	db.save()
	db.delete()
}

/*
Notes:

In Go, interfaces are implemented implicitly.
All you have to do is attach the methods defined in an interface, to a struct,
for that struct to also qualify as the type mentioned by that interface.

- excerpt from "Understanding Interfaces in Go" by Arun Kant Pant
*/
