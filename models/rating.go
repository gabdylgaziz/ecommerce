package models

type Rating struct {
	Id     int  `json:"Id"`
	Value  int  `json:"Value"`
	User   User `json:"User"`
	UserId int  `json:"UserId" gorm:"primaryKey"`
	Item   Item `json:"Item"`
	ItemId int  `json:"ItemId" gorm:"primaryKey"`
}
