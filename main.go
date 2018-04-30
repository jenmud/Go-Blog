package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World")
	t := template.Must(
		template.ParseFiles(
			"views/index.html",
			"views/header.html",
			"views/footer.html"),
	)

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
