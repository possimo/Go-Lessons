package main

import (
	"fmt"
	"lesson14/persondb"
	"strings"
)

var personDb = persondb.InitTable()

var personList = `Ali,30,false
Umi,50,true
Abi,60,true`

func main() {

	person := persondb.Person{} // create empty Person object

	person.New("Ali", 30, false)
	err := personDb.Add("001", person)
	if err != nil {
		fmt.Println(err.Error())
	}

	person.New("Abi", 50, true)
	err = personDb.Add("002", person)
	if err != nil {
		fmt.Println(err.Error())
	}

	personDb.ShowAll()

	person, err = personDb.Get("003")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(person)
	}

	personDb.Delete("002")
	personDb.ShowAll()

	// read person list string from variable
	pList := persondb.PersonList{Reader: strings.NewReader(personList)}
	// initialise person id - integer value
	id := 0
	// retrieve person list into slice - read each person in loop
	for _, person := range pList.Retrieve() {
		// increase person id
		id++
		// format person id in string
		personId := fmt.Sprintf("%03d", id)
		fmt.Println(personId, ":", person)
		// create new person record
		personRec := persondb.Person{}
		// init new person record
		personRec.New(person.Name, person.Age, person.Married)
		// add person record in map database
		err := personDb.Add(personId, personRec)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	// show all person records in map database
	personDb.ShowAll()

}
