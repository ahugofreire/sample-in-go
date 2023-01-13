package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Articles []Article
}

type Article struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:63306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Article{}, &Category{})

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

	//Has Many
	var categories []Category
	err = db.Model(&Category{}).Preload("Articles").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Articles {
			fmt.Println("- ", product.Name, category.Name)
		}

	}
}
