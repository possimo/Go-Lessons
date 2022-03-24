package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"lesson14/persondb"
	"log"
	"net/http"
	"strings"
	"time"
)

const port = "5000"

var personSlice []persondb.Person

func main() {
	// initialise person slice values
	personSlice = initPersonSlice()
	// create a new http servemux/router
	mux := http.NewServeMux()
	// assign a handler to each URL
	mux.HandleFunc("/", indexHandlerText)
	mux.HandleFunc("/personlistjson", personListJsonHandler)
	mux.HandleFunc("/personhtml", personHtmlHandler)
	mux.HandleFunc("/personlisthtml", personListHtmlHandler)
	// set http server config parameters
	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("main: running http server on port", port)
	// start http server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start http server: %v\n", err)
	}
}

func indexHandlerText(w http.ResponseWriter, r *http.Request) {
	// send text response to client browser
	fmt.Fprintf(w, "Hello World")
}

func personListJsonHandler(w http.ResponseWriter, r *http.Request) {
	// convert person slice into JSON bytes string
	personsJSON, err := json.Marshal(personSlice)
	// handle error - send notification to client browser
	if err != nil {
		fmt.Fprintf(w, "Unable to read person list JSON data")
		fmt.Println(err)
		return
	}
	fmt.Println(string(personsJSON))
	// set http response header type
	w.Header().Set("Content-Type", "application/json")
	// set http status value in header
	w.WriteHeader(http.StatusOK)
	// send person list JSON bytes string to requester
	w.Write(personsJSON)
}

func personHtmlHandler(w http.ResponseWriter, r *http.Request) {
	// set http response header type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// read & parse html file
	t, err := template.ParseFiles("templates/person.html")
	// handle error - send notification to client browser
	if err != nil {
		fmt.Fprintf(w, "Unable to load person template")
		fmt.Println(err)
		return
	}
	// get a person from person slice
	person := personSlice[0]
	// send html file & a person to client browser
	t.Execute(w, person)
}

func personListHtmlHandler(w http.ResponseWriter, r *http.Request) {
	// set http response header type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// read & parse html file
	t, err := template.ParseFiles("templates/personlist.html")
	// handle error - send notification to client browser
	if err != nil {
		fmt.Fprintf(w, "Unable to load person list template")
		fmt.Println(err)
		return
	}
	// send html file & person slice to client browser
	t.Execute(w, personSlice)
}

func initPersonSlice() []persondb.Person {
	personList := `Ali,30,false
Umi,50,true
Abi,60,true`

	// read person list string from string variable
	pList := persondb.PersonList{Reader: strings.NewReader(personList)}
	// retrieve person list into slice
	return pList.Retrieve()
}
