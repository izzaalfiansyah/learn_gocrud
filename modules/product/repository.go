package product

import (
	"database/sql"
	"errors"

	"izzaalfiansyah/learn_gocrud/modules/category"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetProducts() (*[]Product, error) {
	query := "SELECT * FROM products"

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Stock, &product.Price, &product.CategoryID)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return &products, nil
}

func (repo *Repository) GetProductByID(id *int) (*Product, error) {
	query := "SELECT p.*, c.name as category_name, c.description as category_description FROM products as p LEFT JOIN categories as c ON c.id = p.category_id WHERE p.id = $1"

	var product Product
	var category category.Category
	err := repo.db.QueryRow(query, &id).Scan(&product.ID, &product.Name, &product.Stock, &product.Price, &product.CategoryID, &category.Name, &category.Description)
	category.ID = product.CategoryID
	product.Category = &category
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *Repository) CreateProduct(newProduct *Product) (*Product, error) {
	query := "INSERT INTO products (name, stock, price, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, &newProduct.Name, &newProduct.Stock, &newProduct.Price, &newProduct.CategoryID).Scan(&newProduct.ID)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (repo *Repository) UpdateProduct(id *int, product *Product) (*Product, error) {
	query := "UPDATE products SET name = $1, stock = $2, price = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, &product.Name, &product.Stock, &product.Price, &product.CategoryID, &id)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("Product not found")
	}

	return product, nil
}

func (repo *Repository) DeleteProduct(id *int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, &id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Product not found")
	}

	return nil
}
