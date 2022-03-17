package main

import (
	"fmt"
	"lesson14/persondb"
	"sync"
	"time"
)

var personSlice []persondb.Person
var personMap = make(map[string]persondb.Person)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	go addLowId()
	addHighId()

	fmt.Println("Person Slice:")
	for idx, person := range personSlice {
		fmt.Printf("%02d : %v\n", idx, person)
	}

	fmt.Println("Person Map:")
	wg.Add(len(personSlice))
	for _, person := range personSlice {
		go func(person persondb.Person, wg *sync.WaitGroup) {
			defer wg.Done()

			mutex.Lock()
			personMap[person.Name] = person
			mutex.Unlock()

			mutex.Lock()
			p := personMap[person.Name]
			mutex.Unlock()

			fmt.Println(p)
		}(person, &wg)
	}
	wg.Wait()
}

func addLowId() {
	for i := 0; i < 10; i++ {
		person := persondb.Person{}
		id := fmt.Sprintf("%03d", i)
		person.New(id, i+10, false)
		personSlice = append(personSlice, person)
	}
}

func addHighId() {
	for i := 99; i > 89; i-- {
		person := persondb.Person{}
		id := fmt.Sprintf("%03d", i)
		person.New(id, i-10, false)
		personSlice = append(personSlice, person)
		time.Sleep(50 * time.Millisecond)
	}
}
