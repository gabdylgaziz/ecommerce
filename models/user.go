package models

type User struct {
	Id       int    `json:"Id" gorm:"primaryKey;autoincrement"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Email    string `json:"Email"`
	Username string `json:"Username"`
	Password string `json:"-"`
}
