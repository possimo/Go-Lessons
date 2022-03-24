package data

import (
	"fmt"
	"lesson19/mysql"
	"strings"
	"time"
)

// Employee type definition
type Employee struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	MobileNo  string `json:"mobileno"`
	Dept      string `json:"dept"`
	BirthDate string `json:"birthdate"`
}

// Add saves new employee record in database
func (e *Employee) Add(db mysql.DB) error {
	// compose SQL command string
	sqlCmd := `INSERT INTO employees (name, email, mobileno, dept,
birthdate, createdat) VALUES (?,?,?,?,?,?)`
	// insert user record in database
	_, err := db.SQL.Exec(sqlCmd, e.Name, e.Email, e.MobileNo, e.Dept,
		reformatDateStr(e.BirthDate), time.Now())
	return err
}

// Retrieve retrieves employee record  from database
func (e *Employee) Retrieve(db mysql.DB) error {
	var day, month, year string
	// compose SQL command string
	sqlCmd := `SELECT name, email, mobileno, dept, DAY(birthdate), 
MONTH(birthdate), YEAR(birthdate) FROM employees WHERE email = ?`
	// get user record from database
	err := db.SQL.QueryRow(sqlCmd, e.Email).Scan(&e.Name, &e.Email,
		&e.MobileNo, &e.Dept, &day, &month, &year)
	e.BirthDate = fmt.Sprintf("%s-%s-%s", day, month, year)
	return err
}

// Update saves employee record changes in database
func (e *Employee) Update(db mysql.DB) error {
	// compose SQL command string
	sqlCmd := "UPDATE employees SET"
	if len(e.Name) > 0 {
		sqlCmd = fmt.Sprintf("%s name = '%s',", sqlCmd, e.Name)
	}
	if len(e.MobileNo) > 0 {
		sqlCmd = fmt.Sprintf("%s mobileno = '%s',", sqlCmd, e.MobileNo)
	}
	if len(e.Dept) > 0 {
		sqlCmd = fmt.Sprintf("%s dept = '%s',", sqlCmd, e.Dept)
	}
	if len(e.BirthDate) > 0 {
		sqlCmd = fmt.Sprintf("%s birthdate = '%s',", sqlCmd, reformatDateStr(e.BirthDate))
	}
	// remove trailing comma from SQL command
	sqlCmd = strings.TrimRight(sqlCmd, ",")
	// add condition statement
	sqlCmd = fmt.Sprintf("%s WHERE email = '%s'", sqlCmd, e.Email)
	// update employee record in database
	_, err := db.SQL.Exec(sqlCmd)
	return err
}

// Delete deletes employee record from database
func (e *Employee) Delete(db mysql.DB) error {
	// get user record from database
	_, err := db.SQL.Exec("DELETE FROM employees WHERE email = ?", e.Email)
	return err
}

// Employees retrieves all employees in the database and returns it
func Employees(db mysql.DB) ([]Employee, error) {
	var empList []Employee
	var day, month, year string
	// compose SQL command string
	sqlCmd := `SELECT name, email, mobileno, dept, DAY(birthdate), 
MONTH(birthdate), YEAR(birthdate) FROM employees`
	// get all employee records from database
	rows, err := db.SQL.Query(sqlCmd)
	if err != nil {
		return empList, err
	}
	defer rows.Close()
	for rows.Next() {
		emp := Employee{}
		if err = rows.Scan(&emp.Name, &emp.Email, &emp.MobileNo, &emp.Dept, &day, &month, &year); err != nil {
			return empList, err
		}
		// compose birth date value
		emp.BirthDate = fmt.Sprintf("%s-%s-%s", day, month, year)
		// append to employees list
		empList = append(empList, emp)
	}
	return empList, err
}

// SearchEmployeeByName searches for employees in database using name substring
func SearchEmployeeByName(name string, db mysql.DB) ([]Employee, error) {
	var empList []Employee
	var day, month, year string
	// compose name search pattern
	findStr := "%" + name + "%"
	// compose SQL command string
	sqlCmd := `SELECT name, email, mobileno, dept, DAY(birthdate), 
MONTH(birthdate), YEAR(birthdate) FROM employees WHERE name LIKE ? ORDER BY name`
	// get all employee records from database with the name pattern
	rows, err := db.SQL.Query(sqlCmd, findStr)
	if err != nil {
		return empList, err
	}
	defer rows.Close()
	for rows.Next() {
		emp := Employee{}
		if err = rows.Scan(&emp.Name, &emp.Email, &emp.MobileNo, &emp.Dept, &day, &month, &year); err != nil {
			return empList, err
		}
		// compose birth date value
		emp.BirthDate = fmt.Sprintf("%s-%s-%s", day, month, year)
		// append to employees list
		empList = append(empList, emp)
	}
	return empList, err
}

// reformats date string from dd-mm-yyyy to yyyy-mm-dd
func reformatDateStr(dateStr string) string {
	// read date fields in day-month-year format
	dateFields := strings.Split(dateStr, "-")
	// reformat date string
	return fmt.Sprintf("%s-%s-%s", dateFields[2], dateFields[1], dateFields[0])
}
