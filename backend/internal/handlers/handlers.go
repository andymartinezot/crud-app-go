package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/andymartinezot/crud-app-go/backend/config"
	"github.com/andymartinezot/crud-app-go/backend/internal/models"
)

var templates *template.Template

func init() {
	// Get the absolute path to the templates directory
	templatesDir := filepath.Join("/app", "internal", "templates", "*")
	// Parse the templates
	templates = template.Must(template.ParseGlob(templatesDir))
}

//var templates = template.Must(template.ParseGlob("../../internal/templates/*"))

func Initiate(w http.ResponseWriter, r *http.Request) {
	setConnection := config.ConnectionDB()
	records, err := setConnection.Query("SELECT * FROM employees")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer records.Close()

	employee := models.Employee{}
	arrayEmployee := []models.Employee{}

	for records.Next() {
		var id int
		var name string
		var email string

		err = records.Scan(&id, &name, &email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrayEmployee = append(arrayEmployee, employee)
	}

	templates.ExecuteTemplate(w, "init", arrayEmployee)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		setConnection := config.ConnectionDB()
		insertRecords, err := setConnection.Prepare("INSERT INTO employees(name,email) VALUES (?,?)")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//insertRecords.Exec(name, email)
		result, err := insertRecords.Exec(name, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("New user added. ID: %v", id)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	setConnection := config.ConnectionDB()
	record, err := setConnection.Query("SELECT * FROM employees where id=?", idEmployee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer record.Close()

	employee := models.Employee{}

	for record.Next() {
		var id int
		var name string
		var email string

		err = record.Scan(&id, &name, &email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employee.Id = id
		employee.Name = name
		employee.Email = email
	}

	templates.ExecuteTemplate(w, "update", employee)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		setConnection := config.ConnectionDB()
		updateRecord, err := setConnection.Prepare("UPDATE employees SET name=?,email=? WHERE id=?")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := updateRecord.Exec(name, email, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rowsAffected > 0 {
			log.Printf("Record with ID %v updated. %v rows affected.", id, rowsAffected)
		} else {
			log.Printf("No record found with ID %v to update.", id)
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	setConnection := config.ConnectionDB()
	deleteRecords, err := setConnection.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := deleteRecords.Exec(idEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 0 {
		log.Printf("Record with ID %v deleted. %v rows affected.", idEmployee, rowsAffected)
	} else {
		log.Printf("No record found with ID %v to delete.", idEmployee)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
