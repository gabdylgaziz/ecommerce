package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h handler) UpdateItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	ErrorHandler(err)

	var item models.Item

	if result := h.DB.Find(&item, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(&item).Updates(item); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("item is updated")
}
