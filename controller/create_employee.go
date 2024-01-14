package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewCreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["name"][0]
			address := r.Form["address"][0]
			npwp := r.Form["npwp"][0]
			_, err := db.Exec("INSERT INTO employee (name, npwp, address) VALUES (?, ?, ?)", name, npwp, address)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
