package handlers

import (
	"log"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
