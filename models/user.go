package models

type User struct {
	Id       int    `json:"Id" gorm:"primaryKey"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Email    string `json:"Email"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	//Orders   []Order   `gorm:"foreignKey:UserId"`
	//Ratings  []Rating  `gorm:"foreignKey:UserId"`
	//Comments []Comment `gorm:"foreignKey:UserId"`
}
