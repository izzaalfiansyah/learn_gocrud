package category

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"izzaalfiansyah/learn_gocrud/config"
	"izzaalfiansyah/learn_gocrud/modules/exception"
)

func CategoryController(w http.ResponseWriter, r *http.Request) {
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

	idStr := strings.ReplaceAll(strings.ReplaceAll(r.URL.Path, "/categories/", ""), "/categories", "")

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
		categories, err := service.GetCategories()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message":    "Categories fetched",
			"categories": &categories,
		})
		return
	}

	if r.Method == "POST" && id == 0 {
		var newCategory Category
		json.NewDecoder(r.Body).Decode(&newCategory)

		category, err := service.AddCategory(&newCategory)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message":  "Category successfully added",
			"category": &category,
		})
		return
	}

	if r.Method == "GET" && id != 0 {
		category, err := service.GetCategoryByID(&id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		if category == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(exception.CreateError(errors.New("Category not found")))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message":  "Category fetched",
			"category": &category,
		})
		return
	}

	if r.Method == "PUT" && id != 0 {
		var updateCategory Category
		json.NewDecoder(r.Body).Decode(&updateCategory)

		category, err := service.UpdateCategory(&id, &updateCategory)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message":  "Category successfully updated",
			"category": &category,
		})
		return
	}

	if r.Method == "DELETE" && id != 0 {
		err := service.DeleteCategory(&id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(exception.CreateError(err))
			return
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Category successfully deleted",
		})
		return
	}
}
