package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h handler) GetItemRating(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	//var rating models.Rating

	value := map[string]interface{}{}

	if result := h.DB.Model(models.Rating{}).Select("AVG(value)").Group("item_id").Having("item_id = ?", id).Take(&value); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)

	fmt.Println("item is printed")
}
