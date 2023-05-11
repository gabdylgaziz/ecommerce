package models

type Rating struct {
	Id     int  `json:"Id" gorm:"autoincrement"`
	Value  int  `json:"Value"`
	User   User `json:"User" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	UserId int  `json:"UserId" gorm:"primaryKey"`
	Item   Item `json:"Item" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	ItemId int  `json:"ItemId" gorm:"primaryKey"`
}
