package main

import (
	//print in console
	"database/sql"
	"fmt"
	"log"           //Console logs
	"net/http"      //WEB server
	"text/template" //Allow use templates from another folder

	_ "github.com/go-sql-driver/mysql" //Import driver to connect with mysql
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {

	http.HandleFunc("/", Initiate)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/delete", Delete)

	log.Println("server up and running...")

	//Initiate server
	http.ListenAndServe(":8000", nil)
}

func connectionDB() (connection *sql.DB) {
	Driver := "mysql"
	User := "root"
	Pass := "password123"
	Name := "system"

	connection, err := sql.Open(Driver, User+":"+Pass+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

// Define structure for the Initiate function
type Employee struct {
	Id    int
	Name  string
	Email string
}

func Initiate(w http.ResponseWriter, r *http.Request) {

	//Read values from the DDBB
	setConnection := connectionDB()
	records, err := setConnection.Query("SELECT * FROM employees")

	if err != nil {
		panic(err.Error())
	}

	employee := Employee{}        //Create var employee which is an structure from the structure Employee
	arrayEmployee := []Employee{} //Create array from the Employee structure

	for records.Next() {
		var id int
		var name string
		var email string

		err = records.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployee = append(arrayEmployee, employee)
	}

	//fmt.Println(arrayEmployee)

	//fmt.Fprintf(w, "Hello from GO!")
	templates.ExecuteTemplate(w, "init", arrayEmployee)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		//Set connection to the DDBB
		setConnection := connectionDB()
		insertRecords, err := setConnection.Prepare("INSERT INTO employees(name,email) VALUES (?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertRecords.Exec(name, email)

		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	//Read values from the DDBB
	setConnection := connectionDB()
	record, err := setConnection.Query("SELECT * FROM employees where id=?", idEmployee)

	employee := Employee{}

	for record.Next() {
		var id int
		var name string
		var email string

		err = record.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email
	}

	fmt.Println(employee)
	templates.ExecuteTemplate(w, "update", employee)

	//http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		//Set connection to the DDBB
		setConnection := connectionDB()
		updateRecord, err := setConnection.Prepare("UPDATE employees SET name=?,email=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		updateRecord.Exec(name, email, id)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	//Set connection to the DDBB
	setConnection := connectionDB()
	deleteRecords, err := setConnection.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	deleteRecords.Exec(idEmployee)

	http.Redirect(w, r, "/", 301)
}
