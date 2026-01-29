package category

import (
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetCategories() (*[]Category, error) {
	query := "SELECT * FROM categories"

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]Category, 0)

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return &categories, nil
}

func (repo *Repository) GetCategoryByID(id *int) (*Category, error) {
	query := "SELECT * FROM categories WHERE id = $1"

	var category Category
	err := repo.db.QueryRow(query, &id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo *Repository) CreateCategory(newCategory *Category) (*Category, error) {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, &newCategory.Name, &newCategory.Description).Scan(&newCategory.ID)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (repo *Repository) UpdateCategory(id *int, category *Category) (*Category, error) {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, &category.Name, &category.Description, &id)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("Category not found")
	}

	return category, nil
}

func (repo *Repository) DeleteCategory(id *int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, &id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Category not found")
	}

	return nil
}
