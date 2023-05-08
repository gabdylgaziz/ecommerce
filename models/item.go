package models

type Item struct {
	Id          int    `json:"Id" gorm:"primaryKey;autoincrement"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
}
