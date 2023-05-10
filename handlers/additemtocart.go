package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h handler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("adding item to cart...")

	id := mux.Vars(r)["id"]

	var item models.Item

	if result := h.DB.Find(&item, id); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userId = claims.Data.Id

	var cart models.Cart

	if result := h.DB.Where("user_id = ?", userId).FirstOrCreate(&cart, models.Cart{UserId: userId}); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(&cart).Association("Items").Append(&item); result != nil {
		fmt.Println(result)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Added successfully")

	fmt.Println("item has been added")

}
