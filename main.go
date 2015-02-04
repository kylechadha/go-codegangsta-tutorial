package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ROUTES
	// -------
	router := mux.NewRouter().StrictSlash(false)
	router.Handle("/", http.FileServer(http.Dir("public")))

	// Posts Collection
	posts := router.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)

	// Posts Singular
	post := router.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	// Markdown
	markdown := router.Path("/markdown").Subrouter()
	markdown.Methods("POST").HandlerFunc(MarkdownHandler)

	// MIDDLEWARE
	// -----------
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(AppMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)
	n.UseHandler(router)

	// SERVER
	// -------
	fmt.Printf("The magic happens on port %s", port)
	http.ListenAndServe(":"+port, router)
	// n.Run(":" + port)
}

func AppMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging in our middleware woot!")

	// This is where you would do cool middleware stuff!
	// if r.URL.Query().Get("password") == "secret123" {
	next(rw, r)
	// } else {
	// http.Error(rw, "Not Authorized", 401)
	// }

	// log.Println("And here it comes back to the middleware")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "post edit")
}

func MarkdownHandler(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
