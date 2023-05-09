package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) PostAddress(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("creating a comment...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var address models.Address
	json.Unmarshal(body, &address)
	address.UserId = claims.Data.Id

	if result := h.DB.Create(&address); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(models.User{}).Where("id = ?", address.UserId).Take(&address.User); result.Error != nil {
		fmt.Println(result.Error)
		h.DB.Where("id = ?", address.Id).Delete(&address)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	fmt.Println("comment has been posted")
}
