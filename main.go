package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Web Developing.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello , Go Web Developing.")
	})

	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}
