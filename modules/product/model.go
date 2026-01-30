package product

import "izzaalfiansyah/learn_gocrud/modules/category"

type Product struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	Stock      string             `json:"stock"`
	Price      string             `json:"price"`
	CategoryID int                `json:"category_id"`
	Category   *category.Category `json:"category"`
}
