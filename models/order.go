package models

type Order struct {
	Id        int     `json:"Id" gorm:"primaryKey;autoincrement"`
	User      User    `json:"User"`
	UserId    int     `json:"UserId"`
	Item      Item    `json:"Item"`
	ItemId    int     `json:"ItemId"`
	Address   Address `json:"Address"`
	AddressId int     `json:"AddressId"`
}
