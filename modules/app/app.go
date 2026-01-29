package app

import (
	"fmt"
	"net/http"

	"izzaalfiansyah/learn_gocrud/config"
	"izzaalfiansyah/learn_gocrud/modules/category"
)

func RunApp() {
	config.LoadConfig()
	port := config.Env.AppPort

	http.HandleFunc("/categories/", category.CategoryController)
	http.HandleFunc("/categories", category.CategoryController)
	http.HandleFunc("/", AppController)

	fmt.Println("Application running on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
