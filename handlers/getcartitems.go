package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) GetCartItems(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("getting items from a cart...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var userItems []models.Item
	var cart models.Cart

	json.Unmarshal(body, &cart)
	var userId = claims.Data.Id

	if result := h.DB.Where("user_id = ?", userId).FirstOrCreate(&cart, models.Cart{UserId: userId}); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(&cart).Association("Items").Find(&userItems); result != nil {
		fmt.Println(result)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userItems)

	fmt.Println("items are sent")

}
