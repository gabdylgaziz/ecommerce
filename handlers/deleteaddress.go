package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("deleting an address...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var address models.Address
	json.Unmarshal(body, &address)

	//if address.UserId != claims.Data.Id {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	json.NewEncoder(w).Encode("You cant do this")
	//	return
	//}

	if result := h.DB.Model(models.Address{}).Where("user_id = ?", claims.Data.Id).Delete(&address); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Deleted")

	fmt.Println("address has been deleted")
}
