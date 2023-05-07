package models

type Order struct {
	Id     int `json:"Id" gorm:"primaryKey"`
	UserId int `json:"UserId"`
	ItemId int `json:"ItemId"`
}
