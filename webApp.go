package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func indexPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to a page about ", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":8000", nil)
}