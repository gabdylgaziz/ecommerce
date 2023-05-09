package db

import (
	"ecommerce/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=dadon2004 dbname=assignmentGo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := "host=localhost user=postgres password=182769 dbname=e_commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to DB")

	db.AutoMigrate(models.User{}, models.Item{}, models.Rating{}, models.Card{}, models.Comment{}, models.Order{}, models.Address{})

	return db
}
