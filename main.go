package main

import (
	"fmt"
	_ "fmt"
	"html/template"
	_ "html/template"
	"net/http"
	_ "net/http"
)

func main() {
	handleFunc()
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
