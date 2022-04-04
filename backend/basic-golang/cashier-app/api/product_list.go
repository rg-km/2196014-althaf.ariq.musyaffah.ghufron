package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)

	allProducts, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(ProductListErrorResponse{Error: err.Error()})
		}
	}()

	products := make([]Product, len(allProducts))
	for i, product := range allProducts {
		products[i] = Product{
			Name:     product.ProductName,
			Price:    product.Price,
			Category: product.Category,
		}
	}

	encoder.Encode(ProductListSuccessResponse{Products: products})
}
