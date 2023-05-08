package models

import "time"

type Card struct {
	Id         int       `json:"Id"`
	Cardholder string    `json:"Cardholder"`
	ExpTime    time.Time `json:"ExpTime"`
	User       User      `json:"User"`
	UserId     int       `json:"UserId" gorm:"primaryKey"`
	Pan        string    `json:"Pan" gorm:"primaryKey;size:16"`
}

// https://stackoverflow.com/questions/60954794/how-to-define-date-in-gorm
