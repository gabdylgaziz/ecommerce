package models

type Address struct {
	Id       int    `json:"Id" gorm:"primaryKey;autoincrement"`
	Country  string `json:"Country"`
	City     string `json:"City"`
	Street   string `json:"Street"`
	Postcode string `json:"Postcode"`
	User     User   `json:"-" gorm:"OnUpdate:CASCADE;OnDelete:CASCADE"`
	UserId   int    `json:"UserId" gorm:"unique"`
}
