package main

import (
	"github.com/russross/blackfriday"
	"net/http"
	"os"
)

func main() {
	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
