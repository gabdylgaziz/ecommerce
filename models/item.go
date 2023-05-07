package models

type Item struct {
	Id          int    `json:"Id" gorm:"primaryKey"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
	//Orders      []Order   `gorm:"foreignKey:ItemId"`
	//Ratings     []Rating  `gorm:"foreignKey:ItemId"`
	//Comments    []Comment `gorm:"foreignKey:ItemId"`
}
