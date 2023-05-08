package models

type Address struct {
	Id       int    `json:"Id" gorm:"primaryKey;autoincrement"`
	Country  string `json:"Country"`
	City     string `json:"City"`
	Street   string `json:"Street"`
	Postcode string `json:"Postcode"`
	User     User   `json:"User"`
	UserId   int    `json:"UserId"`
}
