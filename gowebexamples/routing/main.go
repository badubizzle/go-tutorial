package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func allBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You requested for all books")
}
func getBook(w http.ResponseWriter, r *http.Request) {
	var book = mux.Vars(r)["title"]
	fmt.Fprintf(w, "You requested for book %s\n", book)
}
func getBookPage(w http.ResponseWriter, r *http.Request) {
	var args = mux.Vars(r)
	title := args["title"]
	page := args["page"]

	fmt.Fprintf(w, "You requested for page %s in book %s\n", page, title)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	bookrouter := router.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", allBooks)
	bookrouter.HandleFunc("/{title}", getBook)
	bookrouter.HandleFunc("/{title}/page/{page}", getBookPage)

	http.ListenAndServe(":8090", router)
}
