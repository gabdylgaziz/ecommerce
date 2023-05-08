package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) UpdateRating(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a rating...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var rating models.Rating
	json.Unmarshal(body, &rating)

	h.DB.Where("user_id = ? and item_id = ?", rating.UserId, rating.ItemId).Delete(&rating)

	if result := h.DB.Create(&rating); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(models.User{}).Where("id = ?", rating.UserId).Take(&rating.User); result.Error != nil {
		fmt.Println(result.Error)
		h.DB.Where("id = ?", rating.Id).Delete(&rating)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(models.Item{}).Where("id = ?", rating.ItemId).Take(&rating.Item); result.Error != nil {
		fmt.Println(result.Error)
		h.DB.Where("id = ?", rating.Id).Delete(&rating)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Updated")

	fmt.Println("rating has been updated")

}
