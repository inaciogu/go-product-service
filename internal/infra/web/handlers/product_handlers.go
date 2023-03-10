package handlers

import (
	"encoding/json"
	"net/http"

	usecase "github.com/inaciogu/go-product-service/internal/useCases"
)

type ProductHandlers struct {
	CreateProductUseCase   *usecase.CreateProductUseCase
	ListAllProductsUseCase *usecase.ListAllProductsUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listAllProductsUseCase *usecase.ListAllProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase:   createProductUseCase,
		ListAllProductsUseCase: listAllProductsUseCase,
	}
}

func (h *ProductHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.CreateProductUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(output)
}

func (h *ProductHandlers) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListAllProductsUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
