package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func (h handler) GetMaxPrice() int {
	var item models.Item
	if result := h.DB.Order("price DESC").First(&item); result.Error != nil {
		fmt.Println(result.Error)
		return math.MaxInt
	}
	fmt.Println(item.Price)
	return item.Price
}

func (h handler) GetFilteredItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item

	priceMin := r.URL.Query().Get("min_cost")
	if priceMin == "" {
		priceMin = "0"
	}

	priceMax := r.URL.Query().Get("max_cost")
	if priceMax == "" {
		priceMax = strconv.Itoa(h.GetMaxPrice())
	}

	ratingMin := r.URL.Query().Get("min_rating")
	if ratingMin == "" {
		ratingMin = "0"
	}

	ratingMax := r.URL.Query().Get("max_rating")
	if ratingMax == "" {
		ratingMax = "5"
	}

	if result := h.DB.Where("price >= ? AND price <= ? AND id in (?)", priceMin, priceMax,
		h.DB.Table("ratings").Select("item_id").Group("item_id").Having("AVG(value) >= (?) AND AVG(value) <= (?)",
			ratingMin, ratingMax)).Find(&items); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)

	fmt.Println("item is sent")

}
