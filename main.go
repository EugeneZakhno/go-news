package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handleFunc()
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		return
	}
}
