package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// Create Product godoc
// @Sammary Create product
// @Description Create products
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 500 {object} Error
// @Router /products [post]
// @Security ApiKeyAuth
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

// GetProduct godoc
// @Summay Get Product
// @Description get one product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "prodcut ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [get]
// @Security ApiKeyAuth
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
		errMsg := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errMsg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Sammary Update product
// @Description Update products
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "prodcut ID" Format(uuid)
// @Param request body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 404
// @Failure 500 {object} Error
// @Router /products/{id} [put]
// @Security ApiKeyAuth
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

// DeleteProduct godoc
// @Summay Delete Product
// @Description delete one product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "prodcut ID" Format(uuid)
// @Success 200 {object} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProcuctDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProcuctDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// List Prodcuts godoc
// @Summay List Products
// @Description get all products
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "page number"
// @Param limit query string false "limit"
// @Success 200 {array} entity.Product
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")
	products, err := h.ProcuctDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
