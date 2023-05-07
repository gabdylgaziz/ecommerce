package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (h handler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	if result := h.DB.Find(&items); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (h handler) GetMaxPrice() int {
	var item models.Item
	if result := h.DB.Order("price DESC").First(&item); result.Error != nil {
		fmt.Println(result.Error)
	}
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
		priceMax = string(h.GetMaxPrice())
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
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)

	fmt.Println("item is sent")

}

func (h handler) GetItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	ErrorHandler(err)

	var item models.Item

	if result := h.DB.Find(&item, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("item is printed")
}

func (h handler) CreateItem(w http.ResponseWriter, r *http.Request) {

	fmt.Println("creating an item...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var item models.Item
	json.Unmarshal(body, &item)

	if result := h.DB.Create(&item); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	fmt.Println("item has been created")
}

func (h handler) PostComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating a comment...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var comment models.Comment
	json.Unmarshal(body, &comment)

	fmt.Println(comment)

	comment.CommentDate = time.Now()

	fmt.Println(comment)

	if result := h.DB.Create(&comment); result.Error != nil {
		fmt.Println(result.Error)
	}

	if result := h.DB.Model(models.User{}).Where("id = ?", comment.AuthorId).Take(&comment.Author); result.Error != nil {
		fmt.Println(result.Error)
	}

	if result := h.DB.Model(models.Item{}).Where("id = ?", comment.ItemId).Take(&comment.Item); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	fmt.Println("comment has been posted")
}

func (h handler) PostRating(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating a rating...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var rating models.Rating
	json.Unmarshal(body, &rating)

	if result := h.DB.Create(&rating); result.Error != nil {
		fmt.Println(result.Error)
	}

	if result := h.DB.Model(models.User{}).Where("id = ?", rating.UserId).Take(&rating.User); result.Error != nil {
		fmt.Println(result.Error)
	}

	if result := h.DB.Model(models.Item{}).Where("id = ?", rating.ItemId).Take(&rating.Item); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	fmt.Println("rating has been posted")

}
