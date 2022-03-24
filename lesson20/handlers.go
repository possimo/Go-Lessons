package main

import (
	"fmt"
	"html/template"
	"lesson19/data"
	"log"
	"net/http"
)

const httpRedirect = 302

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// employeeCRUD Create, Read, Update, Delete employee records
func employeeCRUD(writer http.ResponseWriter, request *http.Request) {
	// retrieve all employees
	empList, _ := data.Employees(db)
	// check empty employee list
	empCount := len(empList)
	// create employee list access structure
	var empData struct {
		EmpCount int
		EmpList  []data.Employee
		Label    string
	}
	// assign employee data values
	empData.EmpCount = empCount
	empData.EmpList = empList
	// display employees list page
	generateHTML(writer, empData, "layout", "menu.navbar", "emp.crud")
}

// empSearchByName searches employee by name
func empSearchByName(writer http.ResponseWriter, request *http.Request) {
	// parse employee search form
	if err := request.ParseForm(); err != nil {
		log.Println(fmt.Sprintf("Error: Cannot read name search value - %v", err))
		return
	}
	// get search string value
	searchName := request.PostFormValue("empsearchname")
	// retrieve users list by state
	empList, err := data.SearchEmployeeByName(searchName, db)
	if err != nil {
		log.Println(fmt.Sprintf("Error: Cannot retrieve name search list from database - %v", err))
		return
	}
	// check empty user list
	empCount := len(empList)
	// create user list access structure
	var empData struct {
		EmpCount int
		EmpList  []data.Employee
		Label    string
	}
	// assign employee data values
	empData.EmpCount = empCount
	empData.EmpList = empList
	empData.Label = fmt.Sprintf(" : Find Name '%s'", searchName)
	// display employee search name list page
	generateHTML(writer, empData, "layout", "menu.navbar", "emp.crud")
}

// empAdd create a new employee record
func empAdd(writer http.ResponseWriter, request *http.Request) {
	// parse add new iSurvey profile form
	if err := request.ParseForm(); err != nil {
		log.Println(fmt.Sprintf("Error: Cannot read new employee field values - %v", err))
		return
	}
	// get employee values
	emp := data.Employee{
		Name:      request.PostFormValue("name"),
		Email:     request.PostFormValue("email"),
		MobileNo:  request.PostFormValue("mobileno"),
		Dept:      request.PostFormValue("dept"),
		BirthDate: request.PostFormValue("birthdate"),
	}
	// save new employee record in database
	if err := emp.Add(db); err != nil {
		//msgStr := fmt.Sprintf("Add new Helaian Mata PR user[%s] failed", emp.Email)
		//errorMessage(writer, request, msgStr)
		return
	}
	// return to survey list
	http.Redirect(writer, request, "/", httpRedirect)
}

// empUpdate updates employee record
func empUpdate(writer http.ResponseWriter, request *http.Request) {
	// parse add new iSurvey profile form
	if err := request.ParseForm(); err != nil {
		log.Println(fmt.Sprintf("Error: Cannot read existing employee field values - %v", err))
		return
	}
	// get employee values
	emp := data.Employee{
		Name:      request.PostFormValue("updatename"),
		Email:     request.PostFormValue("updateemail"),
		MobileNo:  request.PostFormValue("updatemobileno"),
		Dept:      request.PostFormValue("updatedept"),
		BirthDate: request.PostFormValue("updatebirthdate"),
	}
	// save new employee record in database
	if err := emp.Update(db); err != nil {
		log.Println(fmt.Sprintf("Error: Employee{%s} update failed - %v", emp.Email, err))
	} else {
		log.Println(fmt.Sprintf("Employee{%s} updated", emp.Email))
	}
	// return to survey list
	http.Redirect(writer, request, "/", httpRedirect)
}

// empDelete deletes employee record
func empDelete(writer http.ResponseWriter, request *http.Request) {
	// parse add new iSurvey profile form
	if err := request.ParseForm(); err != nil {
		log.Println(fmt.Sprintf("Error: Cannot read existing employee field values - %v", err))
		return
	}
	// get user email from user input
	emp := data.Employee{Email: request.PostFormValue("deletemail")}
	// delete employee
	if err := emp.Delete(db); err != nil {
		log.Println(fmt.Sprintf("Error: Employee{%s} delete failed - %v", emp.Email, err))
	} else {
		log.Println(fmt.Sprintf("Employee{%s} deleted", emp.Email))
	}
	// return to survey list
	http.Redirect(writer, request, "/", httpRedirect)
}
