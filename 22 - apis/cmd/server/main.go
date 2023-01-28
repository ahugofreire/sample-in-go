package main

import (
	"log"
	"net/http"

	"github.com/ahugofreire/22/apis/configs"

	"github.com/ahugofreire/22/apis/internal/entity"
	"github.com/ahugofreire/22/apis/internal/infra/database"
	"github.com/ahugofreire/22/apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/ahugofreire/22/apis/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title	Go Expert API Example
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Hugo Freire
// @contact.url github.com/ahugofreire
// @contact.email my@mail.com

// license.name MIT
// license.url  locahost:8080

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
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
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JWTExpiresIn))
	//r.Use(LogRequest) //exemplo de middleware

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate-token", userHandler.GetJWT)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
