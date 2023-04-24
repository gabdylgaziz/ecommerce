package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main_page(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", main_page)
	fmt.Println("server is started")
	http.ListenAndServe(":2004", r)
}
