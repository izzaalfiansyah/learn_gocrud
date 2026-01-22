package category

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"izzaalfiansyah/learn_gocrud/modules/exception"
)

func CategoryController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id int

	idStr := strings.ReplaceAll(r.URL.Path, "/category/", "")

	if idStr != "" {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			json.NewEncoder(w).Encode(exception.CreateError(errors.New("Invalid id")))
			return
		}

		id = idInt
	}

	if r.Method == "GET" && id == 0 {
		categories := GetCategories()
		json.NewEncoder(w).Encode(map[string]any{
			"message":    "Categories fetched",
			"categories": &categories,
		})

		return
	}

	if r.Method == "POST" && id == 0 {
		var newCategory Category
		json.NewDecoder(r.Body).Decode(&newCategory)

		category, err := AddCategory(&newCategory)
		if err != nil {
			json.NewEncoder(w).Encode(exception.CreateError(err))
		}

		json.NewEncoder(w).Encode(map[string]any{
			"message":  "Category successfully added",
			"category": &category,
		})
		return
	}

	if r.Method == "GET" && id != 0 {
		category := GetCategoryById(id)

		json.NewEncoder(w).Encode(map[string]any{
			"message":  "Category fetched",
			"category": &category,
		})
	}
}
