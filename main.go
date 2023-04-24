package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main_page(){
	fmt.Println("Hello world")
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", main_page)
    fmt.Println("server is started")
	http.ListenAndServe(":2004", r)
}
