package main

import "fmt"

type person struct {
	name    string
	age     int
	married bool
}

func (p person) getInfo() string {
	return fmt.Sprintf("%s is %d years old and married status is %t", p.name, p.age, p.married)
}

func (p *person) changeAge(newAge int) {
	p.age = newAge
}

type contact struct {
	code   string
	email  string
	mobile string
}

func (c contact) getInfo() string {
	return fmt.Sprintf("Spy %s email address is %s and mobile number is %s", c.code, c.email, c.mobile)
}

func (c *contact) changeContact(newEmail string, newMobile string) {
	if len(newEmail) > 0 {
		c.email = newEmail
	}
	if len(newMobile) > 0 {
		c.mobile = newMobile
	}
}

type employee struct {
	person
	contact
}

func (e employee) getInfo() string {
	return fmt.Sprintf("%s\n%s", e.person.getInfo(), e.contact.getInfo())
}

func (e *employee) changePerson(newPerson person) {
	// change all field values of the person data
	e.person = newPerson
}

func (e *employee) changeContact(newContact contact) {
	// change code if valid
	if len(newContact.code) > 0 {
		e.contact.code = newContact.code
	}
	// change email if valid
	if len(newContact.email) > 0 {
		e.contact.email = newContact.email
		//e.contact.changeContact(newContact.email, "")
	}
	// // change mobile if valid
	if len(newContact.mobile) > 0 {
		e.contact.mobile = newContact.mobile
		//e.contact.changeContact("", newContact.mobile)
	}
}

func main() {

	ali := person{name: "Ali", age: 20, married: true}
	fmt.Println(ali.getInfo())

	ali.changeAge(30)
	fmt.Println(ali.getInfo())

	spy := contact{code: "001", email: "ejenali@possimo.com", mobile: "0123456789"}
	fmt.Println(spy.getInfo())

	spy.changeContact("ejenali@gmail.com", "")
	fmt.Println(spy.getInfo())

	emp := employee{
		person:  person{name: "Ali", age: 20, married: true},
		contact: contact{code: "001", email: "ejenali@possimo.com", mobile: "0123456789"},
	}
	fmt.Println(emp.getInfo())

	emp.changePerson(person{name: "Ali", age: 40, married: true})
	fmt.Println(emp.getInfo())

	emp.changeContact(contact{email: "ejenali@gmail.com"})
	fmt.Println(emp.getInfo())
}
