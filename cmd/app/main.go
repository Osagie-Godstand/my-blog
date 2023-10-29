package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", handler)

	r.Get("/about", aboutHandler)

	r.Get("/posts", postsHandler)

	r.Get("/project1", project1Handler)

	r.Get("/project2", project2Handler)

	r.Get("/project3", project3Handler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/index.tmpl")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/about.tmpl")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/posts.tmpl")
}

func project1Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/project1/project1.tmpl")
}

func project2Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/project2/project2.tmpl")
}

func project3Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tmpl/project3/project3.tmpl")
}
