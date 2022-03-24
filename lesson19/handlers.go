package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"lesson19/data"
	"log"
	"net/http"
	"strings"
)

type srvResponse struct {
	Msg string
}

func basicAuth(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get basic auth params
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		// check basic auth params
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "user authorization failed", http.StatusUnauthorized)
			return
		}
		// decode basic auth params into ASCII string
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		// read basic auth params values
		pair := strings.SplitN(string(payload), ":", 2)
		// check username & password values
		if len(pair) != 2 || !validParams(pair[0], pair[1]) {
			http.Error(w, "Error: User authorization failed", http.StatusUnauthorized)
			return
		}
		// pass control to the next function
		next.ServeHTTP(w, r)
	})
}

func validParams(name, pswd string) bool {
	// compare basic auth params with values
	if name == username && pswd == password {
		return true
	}
	return false
}

func readEmployeeHttpReq(w http.ResponseWriter, r *http.Request) (data.Employee, error) {
	var employee data.Employee
	// read request body
	reqBody, err := io.ReadAll(r.Body)
	// handle error in reading http message body
	if err != nil {
		// compose error message
		errMsg := fmt.Sprintf("Error: Cannot read http message body: %v", err)
		// log error message
		log.Println(errMsg)
		// send error message to client
		http.Error(w, errMsg, http.StatusBadRequest)
		// close http request body
		defer r.Body.Close()
		// return empty employee request data & error
		return employee, err
	}
	// close http request body later if no read error
	defer r.Body.Close()
	// get request params data in JSON format
	if err = json.Unmarshal(reqBody, &employee); err != nil {
		// compose error message
		errMsg := fmt.Sprintf("Error: Cannot read JSON request parameters: %v", err)
		// log error message
		log.Println(errMsg)
		// send error message to client
		http.Error(w, errMsg, http.StatusBadRequest)
	}
	// return employee request data & error
	return employee, err
}

func sendJsonHttpResp(w http.ResponseWriter, httpStatus int, data interface{}) {
	// convert data in json byte string format
	jsonBytes, _ := json.Marshal(data)
	// set http response header type
	w.Header().Set("Content-Type", "application/json")
	// set http status value in header
	w.WriteHeader(httpStatus)
	// send person list JSON bytes string to requester
	w.Write(jsonBytes)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	// read client request for employee data
	emp, err := readEmployeeHttpReq(w, r)
	if err != nil {
		return
	}
	// add employee record in database
	if err = emp.Add(db); err != nil {
		// compose error message
		errMsg := fmt.Sprintf("Error: Cannot add employee record to database: %v", err)
		// log error message
		log.Println(errMsg)
		// send error message to client
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	// send response message to client
	okMsg := fmt.Sprintf("Employee{%s} added to database", emp.Email)
	log.Println(okMsg)
	sendJsonHttpResp(w, http.StatusOK, srvResponse{okMsg})
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	// read client request for employee data
	emp, err := readEmployeeHttpReq(w, r)
	if err != nil {
		return
	}
	// retrieve employee record in database
	if err = emp.Retrieve(db); err != nil {
		// compose error message
		errMsg := fmt.Sprintf("Error: Cannot retrieve employee record from database: %v", err)
		// log error message
		log.Println(errMsg)
		// send error message to client
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	// send employee data to client
	log.Println(emp)
	sendJsonHttpResp(w, http.StatusOK, emp)
}

func updEmployee(w http.ResponseWriter, r *http.Request) {
	// read client request for employee data
	emp, err := readEmployeeHttpReq(w, r)
	if err != nil {
		return
	}
	// update employee record in database
	if err = emp.Update(db); err != nil {
		// compose error message
		errMsg := fmt.Sprintf("Error: Cannot update employee record in database: %v", err)
		// log error message
		log.Println(errMsg)
		// send error message to client
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	// send response message to client
	okMsg := fmt.Sprintf("Employee{%s} updated in database", emp.Email)
	log.Println(okMsg)
	sendJsonHttpResp(w, http.StatusOK, srvResponse{okMsg})
}

func delEmployee(w http.ResponseWriter, r *http.Request) {
	// read client request for employee data
	emp, err := readEmployeeHttpReq(w, r)
	if err != nil {
		return
	}
	// delete employee record in database
	if err = emp.Delete(db); err != nil {
		// compose error message
		errMsg := fmt.Sprintf("Error: Cannot delete employee record in database: %v", err)
		// log error message
		log.Println(errMsg)
		// send error message to client
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	// send response message to client
	okMsg := fmt.Sprintf("Employee{%s} deleted from database", emp.Email)
	log.Println(okMsg)
	sendJsonHttpResp(w, http.StatusOK, srvResponse{okMsg})
}
