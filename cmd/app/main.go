package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	r.Get("/project4", project4Handler)

	r.Get("/blog1", blog1Handler)

	r.Get("/blog2", blog2Handler)

	r.Get("/blog3", blog3Handler)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/index.tmpl")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/about.tmpl")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/posts.tmpl")
}

func project1Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/project/project1/project1.tmpl")
}

func project2Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/project/project2/project2.tmpl")
}

func project3Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/project/project3/project3.tmpl")
}

func project4Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/project/project4/project4.tmpl")
}

func blog1Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/blog/blog1/blog1.tmpl")
}

func blog2Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/blog/blog2/blog2.tmpl")
}

func blog3Handler(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "tmpl/blog/blog3/blog3.gohtml")
}

func serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.Open(filePath)
	checkError(w, err)
	defer file.Close()

	fi, err := file.Stat()
	checkError(w, err)

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)
}

func checkError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error opening file:", err)
		return
	}
}
