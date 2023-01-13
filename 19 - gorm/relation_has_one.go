package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Article struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey`
	Number    string
	ArticleID int
}

func main() {
	dsn := "root:root@tcp(localhost:63306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Article{}, &Category{}, &SerialNumber{})

	//Create Category
	// category := Category{Name: "outros"}
	// db.Create(&category)

	// //Create Product
	// db.Create(&Article{
	// 	Name:       "Adaptador",
	// 	Price:      135.65,
	// 	CategoryID: 2,
	// })

	// db.Create(&SerialNumber{
	// 	Number:    "123465",
	// 	ArticleID: 6,
	// })

	//Has One Relation
	var products []Article
	db.Preload("category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
