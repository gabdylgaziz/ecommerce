package main

import (
	"ecommerce/packages"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
func main_page(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", main_page)
	r.HandleFunc("/signin", packages.Signin)
	r.HandleFunc("/welcome", packages.Welcome)
	r.HandleFunc("/refresh", packages.Refresh)
	r.HandleFunc("/logout", packages.Logout)
	fmt.Println("server is started")
	http.ListenAndServe(":2004", r)
}
