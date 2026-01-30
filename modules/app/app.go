package app

import (
	"fmt"
	"net/http"

	"izzaalfiansyah/learn_gocrud/config"
	"izzaalfiansyah/learn_gocrud/modules/category"
	"izzaalfiansyah/learn_gocrud/modules/check_health"
	"izzaalfiansyah/learn_gocrud/modules/product"
)

func RunApp() {
	config.LoadConfig()
	port := config.Env.AppPort

	http.HandleFunc("/check-health", check_health.CheckHealthController)
	http.HandleFunc("/categories/", category.CategoryController)
	http.HandleFunc("/categories", category.CategoryController)
	http.HandleFunc("/products/", product.ProductController)
	http.HandleFunc("/products", product.ProductController)
	http.HandleFunc("/", AppController)

	fmt.Println("Application running on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
