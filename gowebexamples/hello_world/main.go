package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested %s\n", r.URL.Path)
}
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8090", nil)
}
