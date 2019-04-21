package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name     string
	DBstatus bool
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	// fmt.Println("Hello, Go Web Developing.")

	db, _ := sql.Open("sqlite3", "dev.// DEBUG: ")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Gopher"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// fmt.Fprint(w, "Hello , Go Web Developing.")
	})

	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}
