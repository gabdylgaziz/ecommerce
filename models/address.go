package models

type Address struct {
	Id       int    `json:"Id" gorm:"autoincrement"`
	Country  string `json:"Country"`
	City     string `json:"City"`
	Street   string `json:"Street"`
	Postcode string `json:"Postcode"`
	User     User   `json:"User"`
	UserId   int    `json:"UserId" gorm:"primaryKey"`
}
