package main

import (
	"net/http"

	"github.com/ahugofreire/22/apis/configs"

	"github.com/ahugofreire/22/apis/internal/entity"
	"github.com/ahugofreire/22/apis/internal/infra/database"
	"github.com/ahugofreire/22/apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	http.ListenAndServe(":8000", r)

}
