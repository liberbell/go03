package main

import (
	"fmt"
	"html/template"
	"net/http"
)

Type Page struct {
  Name string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	// fmt.Println("Hello, Go Web Developing.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// fmt.Fprint(w, "Hello , Go Web Developing.")
	})

	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}
