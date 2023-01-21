package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahugofreire/22/apis/internal/dto"
	"github.com/ahugofreire/22/apis/internal/entity"
	"github.com/ahugofreire/22/apis/internal/infra/database"
	entityPkg "github.com/ahugofreire/22/apis/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProcuctDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProcuctDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProcuctDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println(id, "aqui é o id ")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProcuctDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println(id, "aqui é o id ")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProcuctDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProcuctDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
