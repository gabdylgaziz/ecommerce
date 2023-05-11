package models

type Order struct {
	Id        int     `json:"Id" gorm:"primaryKey;autoincrement"`
	User      User    `json:"User" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	UserId    int     `json:"UserId"`
	Item      Item    `json:"Item" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	ItemId    int     `json:"ItemId"`
	Address   Address `json:"Address" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	AddressId int     `json:"AddressId" gorm:"foreignKey"`
}
