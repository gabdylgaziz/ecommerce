package handlers

import (
	"ecommerce/packages"
	"fmt"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

func getData(c *http.Cookie) packages.Claims {
	tknStr := c.Value
	claims := &packages.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("error")
		}
		fmt.Println("error")
	}
	if !tkn.Valid {
		fmt.Println("error")
	}

	return *claims
}