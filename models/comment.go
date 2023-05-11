package models

import "time"

type Comment struct {
	Id          int       `json:"Id" gorm:"primaryKey;autoincrement"`
	Text        string    `json:"Text"`
	Author      User      `json:"Author" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	AuthorId    int       `json:"AuthorId" gorm:"not null"`
	Item        Item      `json:"Item" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	ItemId      int       `json:"ItemId" gorm:"not null"`
	CommentDate time.Time `json:"CommentDate"`
}

//https://www.mindbowser.com/golang-go-with-gorm/
