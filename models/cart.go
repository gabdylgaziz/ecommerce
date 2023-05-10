package models

type Cart struct {
	//Id     int    `json:"Id" gorm:"primaryKey;autoincrement"`
	User   User   `json:"User"`
	UserId int    `json:"UserId" gorm:"primaryKey;"`
	Items  []Item `json:"Item" gorm:"many2many:cart_items;foreignKey:UserId"`
}
