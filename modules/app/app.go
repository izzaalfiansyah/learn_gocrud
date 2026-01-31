package app

import (
	"fmt"
	"log"
	"net/http"

	"izzaalfiansyah/learn_gocrud/config"
	"izzaalfiansyah/learn_gocrud/modules/category"
	"izzaalfiansyah/learn_gocrud/modules/check_health"
	"izzaalfiansyah/learn_gocrud/modules/product"
)

func RunApp() {
	config.LoadConfig()
	_, err := config.InitDB()
	if err != nil {
		log.Println("Failed connect to database:", err.Error())
	}

	port := config.Env.AppPort

	http.HandleFunc("/check-health", check_health.CheckHealthController)
	http.HandleFunc("/categories/", category.CategoryController)
	http.HandleFunc("/categories", category.CategoryController)
	http.HandleFunc("/products/", product.ProductController)
	http.HandleFunc("/products", product.ProductController)
	http.HandleFunc("/", AppController)

	log.Println("Application running on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
