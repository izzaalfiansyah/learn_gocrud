package product

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"izzaalfiansyah/learn_gocrud/config"
	"izzaalfiansyah/learn_gocrud/modules/exception"
)

func ProductController(w http.ResponseWriter, r *http.Request) {
	db, err := config.DB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(exception.CreateError(err))
		return
	}

	repo := NewRepository(db)
	service := NewService(repo)

	w.Header().Set("Content-Type", "application/json")
	var id int

	idStr := strings.ReplaceAll(strings.ReplaceAll(r.URL.Path, "/products/", ""), "/products", "")

	if idStr != "" {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		id = idInt
	}

	if r.Method == "GET" && id == 0 {
		products, err := service.GetProducts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message":  "Products fetched",
			"products": &products,
		})
		return
	}

	if r.Method == "POST" && id == 0 {
		var newProduct Product
		json.NewDecoder(r.Body).Decode(&newProduct)

		product, err := service.AddProduct(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Product successfully added",
			"product": &product,
		})
		return
	}

	if r.Method == "GET" && id != 0 {
		product, err := service.GetProductByID(&id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(exception.CreateError(errors.New("Product not found")))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Product fetched",
			"product": &product,
		})
		return
	}

	if r.Method == "PUT" && id != 0 {
		var updateProduct Product
		json.NewDecoder(r.Body).Decode(&updateProduct)

		product, err := service.UpdateProduct(&id, &updateProduct)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Product successfully updated",
			"product": &product,
		})
		return
	}

	if r.Method == "DELETE" && id != 0 {
		err := service.DeleteProduct(&id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Product successfully deleted",
		})
		return
	}
}
