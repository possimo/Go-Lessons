package persondb

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Person struct {
	Name      string
	Age       int
	Married   bool
	createdAt time.Time // local field
}

type PersonDb struct {
	mutex sync.Mutex
	table map[string]Person
}

type dbError struct {
	code    string
	message string
}

func InitTable() *PersonDb {
	return &PersonDb{
		table: make(map[string]Person),
	}
}

// complies with Stringer interface
func (p Person) String() string {
	// compose createdAt time string
	timeStr := fmt.Sprintf("%d.%02d.%02dT%02d:%02d:%02d",
		p.createdAt.Year(), p.createdAt.Month(), p.createdAt.Day(),
		p.createdAt.Hour(), p.createdAt.Minute(), p.createdAt.Second())
	// return formatted Person data
	return fmt.Sprintf("Person{Name: %s  Age: %d  Married: %t   Created: %s}",
		p.Name, p.Age, p.Married, timeStr)
}

// complies with Error interface
func (e *dbError) Error() string {
	return fmt.Sprintf("Error %s: %s", e.code, e.message)
}

func (p *Person) New(name string, age int, married bool) {
	p.Name = name
	p.Age = age
	p.Married = married
	p.createdAt = time.Now()
}

func (pdb *PersonDb) Add(key string, person Person) error {
	var err error
	pdb.mutex.Lock()
	// check if the key already exist
	if _, ok := pdb.table[key]; ok {
		err = &dbError{code: "901", message: fmt.Sprintf("Add failed - Person key{%s} already recorded in database", key)}
	} else {
		// if a new person - add to database
		pdb.table[key] = person
	}
	pdb.mutex.Unlock()
	return err
}

func (pdb *PersonDb) Get(key string) (Person, error) {
	var err error
	pdb.mutex.Lock()
	person, ok := pdb.table[key]
	if !ok {
		err = &dbError{code: "902", message: fmt.Sprintf("Person key{%s} not found in database", key)}
	}
	pdb.mutex.Unlock()
	return person, err
}

func (pdb *PersonDb) Delete(key string) error {
	var err error
	pdb.mutex.Lock()
	_, ok := pdb.table[key]
	if !ok {
		err = &dbError{code: "903", message: fmt.Sprintf("Delete failed - Person key{%s} not found in database", key)}
	} else {
		delete(pdb.table, key)
	}
	pdb.mutex.Unlock()
	return err
}

func (pdb *PersonDb) ShowAll() {
	pdb.mutex.Lock()
	for _, person := range pdb.table {
		fmt.Println(person)
	}
	pdb.mutex.Unlock()
}

// create type for reading person list from string/text
type PersonList struct {
	Reader io.Reader
}

func (pl PersonList) Retrieve() []Person {
	var pList []Person
	fmt.Println("Initial Person Slice Length:", len(pList))
	// create bytes slice buffer for reading person data in string - 1Kbytes size
	buffer := make([]byte, 1024)
	// read string data into bytes slice buffer
	nbytes, _ := pl.Reader.Read(buffer)
	// convert bytes slice buffer into ASCII string
	personString := string(buffer[:nbytes])
	// create a new slice containing each line of person data - split using new line char
	personLine := strings.Split(personString, "\n")
	// read each line of person data in the slice
	for _, p := range personLine {
		// create a new slice containing each field of person data - split using comma
		pField := strings.Split(p, ",")
		fmt.Println("Person Field Slice", pField, "Length:", len(pField))
		// read only valid person data - with 3 fields
		if len(pField) > 1 {
			// get person field value from each slice element using index
			name := pField[0]                          // assign direct string value
			age, _ := strconv.Atoi(pField[1])          // convert string to integer
			married, _ := strconv.ParseBool(pField[2]) // convert string to boolean
			// create a new person object & assign values read from the person field slice
			person := Person{Name: name, Age: age, Married: married}
			// append person to the person list
			pList = append(pList, person)
		}
	}
	return pList
}
