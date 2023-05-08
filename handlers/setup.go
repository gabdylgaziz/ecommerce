package handlers

import (
	"ecommerce/db"
	"ecommerce/packages"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

var h = New(db.Connect())
var r = mux.NewRouter()

//var itemsR = r.PathPrefix("/items").Subrouter()

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func HandleRequests() {

	r.HandleFunc("/", mainPage)
	r.HandleFunc("/signin", packages.Signin)
	r.HandleFunc("/welcome", packages.Welcome)
	r.HandleFunc("/refresh", packages.Refresh)
	r.HandleFunc("/logout", packages.Logout)

	r.HandleFunc("/items/all", h.GetAllItems).Methods("GET")

	r.Path("/items").Queries("id", "{id}").HandlerFunc(h.GetItemById).Methods("GET")
	r.HandleFunc("/items/{id}/rating", h.PostRating).Methods("POST")
	r.HandleFunc("/items/{id}/rating", h.UpdateRating).Methods("PUT")
	r.HandleFunc("/items/{id}/comment", h.PostComment).Methods("POST")
	r.HandleFunc("/items", h.GetFilteredItems).Methods("GET")
	r.HandleFunc("/items", h.CreateItem).Methods("POST")

	fmt.Println("server is started")
	http.ListenAndServe(":2004", r)
}
