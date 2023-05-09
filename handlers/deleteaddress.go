package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting an address...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var address models.Address
	json.Unmarshal(body, &address)

	if result := h.DB.Model(models.Address{}).Where("id = ?", address.Id).Delete(&address); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Updated")

	fmt.Println("address has been updated")
}
