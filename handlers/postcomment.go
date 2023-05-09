package handlers

import (
	"ecommerce/models"
	//"ecommerce/packages"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	//"github.com/golang-jwt/jwt/v4"
)



func (h handler) PostComment(w http.ResponseWriter, r *http.Request) {
	//чтобы взять данные об авторизованном пользователе
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
	//fmt.Println(claims)

	fmt.Println("creating a comment...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var comment models.Comment
	json.Unmarshal(body, &comment)

	//fmt.Println(claims.Data.Id)
	//fmt.Println(comment)
	
	//присвоить айдишку автора к комменту
	comment.AuthorId = claims.Data.Id

	comment.CommentDate = time.Now()

	if result := h.DB.Create(&comment); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(models.User{}).Where("id = ?", comment.AuthorId).Take(&comment.Author); result.Error != nil {
		fmt.Println(result.Error)
		h.DB.Where("id = ?", comment.Id).Delete(&comment)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if result := h.DB.Model(models.Item{}).Where("id = ?", comment.ItemId).Take(&comment.Item); result.Error != nil {
		fmt.Println(result.Error)
		h.DB.Where("id = ?", comment.Id).Delete(&comment)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	fmt.Println("comment has been posted")
}
