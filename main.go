package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

func create(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	t.ExecuteTemplate(w, "create", nil)
}

var posts = []Article{}
var showPosts = []Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html"))

	db, err := sql.Open("postgres", "postgresql://godbtest_user:lUDEQDsf2MrpRu80RajTBSOG70RNBcY4@dpg-cu74g1q3esus73fg1beg-a.oregon-postgres.render.com/godbtest_21mb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Выборка данных
	res, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}

	//posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)

		fmt.Println(fmt.Sprintf("Post: %s with id %d", post.Title, post.Id))
	}
	t.ExecuteTemplate(w, "index", posts)

}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full_text")

	if title == "" || anons == "" || fullText == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, err := sql.Open("postgres", "postgresql://godbtest_user:lUDEQDsf2MrpRu80RajTBSOG70RNBcY4@dpg-cu74g1q3esus73fg1beg-a.oregon-postgres.render.com/godbtest_21mb")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		//Установка данных:
		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (title, anons,full_text) VALUES ('%s', '%s','%s')", title, anons, fullText))
		if err != nil {
			panic(err)
		}
		defer insert.Close()
		fmt.Println("Успешно добавлено!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func showPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := sql.Open("postgres", "postgresql://godbtest_user:lUDEQDsf2MrpRu80RajTBSOG70RNBcY4@dpg-cu74g1q3esus73fg1beg-a.oregon-postgres.render.com/godbtest_21mb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(vars)
	// выборка данных

	//// Выборка данных
	//res, err := db.Query("SELECT * FROM articles")
	//if err != nil {
	//	panic(err)
	//}
	//for res.Next() {
	//	var user User
	//	err res. Scan(&user.Name, &user.Age)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(fmt.Sprintf("User: %s with age %d", user. Name, user.Age))
	//}
	//t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", saveArticle).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", create).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("./static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
