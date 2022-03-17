package main

import (
	"fmt"
	"lesson14/persondb"
	"time"
)

type dbRequest struct {
	ops   string
	key   string
	rec   persondb.Person
	chRes chan dbResponse
}

type dbResponse struct {
	dbResult string
}

var personDb = persondb.InitTable()

func main() {
	// create specific channels by types
	chReq := make(chan dbRequest)
	chRes := make(chan dbResponse)
	chQuit := make(chan struct{})

	// start database server as a goroutine - listening indefinitely
	go dbServer(chReq, chQuit)

	person := persondb.Person{}
	person.New("Ali", 30, false)

	dbReq := dbRequest{ops: "add", key: "001", rec: person, chRes: chRes}
	dbSendRequest(chReq, dbReq)

	dbReq.ops = "get"
	dbReq.key = "002"
	dbSendRequest(chReq, dbReq)

	dbReq.ops = "del"
	dbReq.key = "001"
	dbSendRequest(chReq, dbReq)

	personDb.ShowAll()

	// terminate database server process
	chQuit <- struct{}{}
}

func dbServer(chReq chan dbRequest, chQuit chan struct{}) {
	// initialise database response
	response := dbResponse{}
	// forever handles request from database client
	for {
		select {
		case request := <-chReq:
			switch request.ops {
			case "add":
				err := personDb.Add(request.key, request.rec)
				if err != nil {
					response.dbResult = err.Error()
				} else {
					response.dbResult = "Add OK"
				}
				request.chRes <- response
			case "get":
				person, err := personDb.Get(request.key)
				if err != nil {
					response.dbResult = err.Error()
				} else {
					response.dbResult = fmt.Sprint(person)
				}
				request.chRes <- response
			case "del":
				err := personDb.Delete(request.key)
				if err != nil {
					response.dbResult = err.Error()
				} else {
					response.dbResult = "Delete OK"
				}
				request.chRes <- response
			}
		case <-chQuit:
			fmt.Println("Database server terminating...")
			time.Sleep(50 * time.Millisecond)
			return
		}
	}
}

func dbSendRequest(chReq chan dbRequest, request dbRequest) {
	chReq <- request
	dbRes := <-request.chRes
	fmt.Println(dbRes)
}
