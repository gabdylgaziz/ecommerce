package models

import "time"

type Comment struct {
	Id          int       `json:"Id" gorm:"primaryKey"`
	Text        string    `json:"Text"`
	Author      User      `json:"Author"`
	AuthorId    int       `json:"AuthorId" gorm:"not null"`
	Item        Item      `json:"Item"`
	ItemId      int       `json:"ItemId" gorm:"not null"`
	CommentDate time.Time `json:"CommentDate"`
}

//https://www.mindbowser.com/golang-go-with-gorm/
