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

	//RELACIONAMENTO
	//Create category
	category := Category{
		Name: "Eletronicos",
	}
	db.Create(&category)

	//create article
	db.Create(&Article{
		Name:       "Notebook",
		Price:      4500.00,
		CategoryID: category.ID,
	})

	//Busca realcionamento
	var products []Article
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}

	//Create
	db.Create(&Article{
		Name:  "Iphone",
		Price: 5000.0,
	})

	//Create in batch
	articles := []Article{
		{Name: "Galaxy 10", Price: 2500.0},
		{Name: "Mouse", Price: 164.54},
		{Name: "MotoG7", Price: 980.56},
	}
	db.Create(&articles)

	//Select One
	var article Article
	db.First(&article, 3)
	fmt.Println(article)

	//Select One with WHERE
	db.First(&article, "name = ?", "Iphone")
	fmt.Println(article)

	//Select all
	// var articles []Article
	db.Find(&articles)
	for _, article := range articles {
		fmt.Println(article)
	}

	//Select all with limit and offset
	db.Limit(2).Offset(2).Find(&articles)
	for _, article := range articles {
		fmt.Println(article)
	}

	//Where
	db.Where("price > ?", 1000).Find(&articles)
	for _, article := range articles {
		fmt.Println(article)
	}

	//Where with like
	db.Where("name LIKE ?", "Mo%").Find(&articles)
	for _, article := range articles {
		fmt.Println(article)
	}

	//Update
	var a Article
	db.First(&a, 3)
	a.Name = "Capinha do iphone"
	db.Save(&a)

	//Delete
	var a2 Article
	db.First(&a2, 3)
	db.Delete(&a2)

}
