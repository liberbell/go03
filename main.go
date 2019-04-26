package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name     string
	DBstatus bool
}

type SearchResult struct {
	Title  string 'xml:"title,attr"'
	Author string 'xml:"author,attr"'
	Year   string 'xml:"hyr,attr"'
	ID     string 'xml:"owi,attr"'
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	// fmt.Println("Hello, Go Web Developing.")

	db, _ := sql.Open("sqlite3", "dev.db")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Gopher"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		p.DBstatus = db.Ping() == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// fmt.Fprint(w, "Hello , Go Web Developing.")
		// db.Close()
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results := []SearchResult{
			SearchResult{"Mody Dick", "Helman Melville", "1851", "222222"},
			SearchResult{"The Adventures of Huckleverry Finn", "Mark Twain", "1884", "444444"},
			SearchResult{"The Cather in the Ray", "JD Salinger", "1951", "333333"},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

type ClassifySearchRespnse struct {
  Results []SearchResult 'xml: "works>work"'

}

func search(query string) ([]SearchResult, error) {
	var resp *http.Response
	var err error
	if resp, err = http.Get("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query)); err != nil {
		return []SearchResult{}, err
	}

	defer resp.Body.Close()
  bar Body []byte
  if Body, err = ioutil.ReadAll(resp.Body); err != nil {
    return []SearchResult(), err
  }
}
