package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World")
	t, err := template.ParseFiles(
		"./views/index.html",
		"./views/header.html",
		"./views/footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	return
	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
