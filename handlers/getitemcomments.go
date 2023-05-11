package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h handler) GetItemComments(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var comments []models.Comment

	if result := h.DB.Where("item_id = ?", id).Preload("Author").
		Preload("Item").Find(&comments); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)

	fmt.Println("item is printed")
}
