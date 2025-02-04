package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func create(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	t.ExecuteTemplate(w, "create", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	t.ExecuteTemplate(w, "index", nil)
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("title")
	fmt.Println(title, anons, fullText)
}

func handleFunc() {
	http.Handle("./static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", saveArticle)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
