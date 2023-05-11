package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Please authorize")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims := getData(c)

	fmt.Println("getting orders for a user...")

	var id = claims.Data.Id
	var orders []models.Order

	if result := h.DB.Where("user_id = ?", id).Preload("User").Preload("Item").Preload("Address").Find(&orders); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)

	fmt.Println("item is printed")
}
