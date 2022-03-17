package main

import (
	"fmt"
	"lesson14/persondb"
	"sync"
)

var personSlice []persondb.Person

var personDb = persondb.InitTable()

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	//go addLowId()
	go addLowIdWg(&wg)
	addHighId()

	wg.Wait()

	fmt.Println("Person Slice:")
	for idx, person := range personSlice {
		fmt.Printf("%02d : %v\n", idx, person)
	}

	fmt.Println("Person Map Database:")
	wg.Add(len(personSlice))
	for _, person := range personSlice {
		go func(person persondb.Person, wg *sync.WaitGroup) {
			personDb.Add(person.Name, person)
			wg.Done()
		}(person, &wg)
	}
	wg.Wait()
	personDb.ShowAll()

}

func addLowId() {
	for i := 0; i < 10; i++ {
		person := persondb.Person{}
		id := fmt.Sprintf("%03d", i)
		person.New(id, i+10, false)
		personSlice = append(personSlice, person)
	}
}

func addLowIdWg(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		person := persondb.Person{}
		id := fmt.Sprintf("%03d", i)
		person.New(id, i+10, false)
		personSlice = append(personSlice, person)
	}
	wg.Done()
}

func addHighId() {
	for i := 99; i > 89; i-- {
		person := persondb.Person{}
		id := fmt.Sprintf("%03d", i)
		person.New(id, i-10, false)
		personSlice = append(personSlice, person)
		//time.Sleep(50 * time.Millisecond)
	}
}
