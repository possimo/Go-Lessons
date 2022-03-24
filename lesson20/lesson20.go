package main

import (
	"fmt"
	"lesson19/mysql"
	"log"
	"net/http"
	"time"
)

const port = "5000"

var db mysql.DB

func main() {
	// connect to MySQL server
	err := db.ConnectMySQL()
	if err != nil {
		log.Fatalf("main: Connect to MySQL database server failed: %v\n", err)
	}
	// close database connection when the program terminates
	defer db.SQL.Close()

	// create a new http servemux/router
	mux := http.NewServeMux()

	// handle static assets
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// assign a handler to each URL
	mux.HandleFunc("/", employeeCRUD)
	mux.HandleFunc("/empsearchbyname", empSearchByName)
	mux.HandleFunc("/empadd", empAdd)
	mux.HandleFunc("/empupdate", empUpdate)
	mux.HandleFunc("/empdelete", empDelete)

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
