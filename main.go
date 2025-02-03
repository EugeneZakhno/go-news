package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
	}
	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
	}
	t.ExecuteTemplate(w, "create", nil)
}

func handleFunc() {
	http.Handle("./static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
