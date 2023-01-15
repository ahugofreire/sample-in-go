package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:63306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Lock  otimista
	//trabalha com versionamento, quando esta alterando criar uma versão

	//Lock Pessimista
	//Locka a table, a linha que vc está alterando do db e naquele momento ninguem consegue alterar.

	tx := db.Begin{}
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Outros"
	tx.Debug().Save(&c)
	tx.Commit()

	// db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// category2 := Category{Name: "Casa"}
	// db.Create(&category2)

	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      99.00,
	// 	Categories: []Category{category, category2},
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("_", product.Name)
		}
	}
}
