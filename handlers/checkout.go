package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) Checkout(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("checking out...")

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

	total := map[string]interface{}{}

	if result := h.DB.Table("cart_items").Select("cart_items.cart_user_id, SUM(price)").
		Joins("inner join items on cart_items.item_id = items.id").
		Group("cart_items.cart_user_id").Having("cart_user_id = ?", userId).Take(&total); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(&cart).Association("Items").Find(&userItems); result != nil {
		fmt.Println(result)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, item := range userItems {
		var user models.User
		var address models.Address

		if result := h.DB.Where("id = ?", userId).First(&user); result.Error != nil {
			fmt.Println(result.Error)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if result := h.DB.Where("user_id = ?", userId).First(&address); result.Error != nil {
			fmt.Println(result.Error)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("passed")

		var order = models.Order{
			User:      user,
			UserId:    userId,
			Item:      item,
			ItemId:    item.Id,
			Address:   address,
			AddressId: address.Id,
		}

		if result := h.DB.Create(&order); result.Error != nil {
			fmt.Println(result.Error)
		}

		fmt.Println("passed")
	}

	if result := h.DB.Model(&cart).Association("Items").Delete(&userItems); result != nil {
		fmt.Println(result)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("passed")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(total)

	fmt.Println("items are checked out")

}
