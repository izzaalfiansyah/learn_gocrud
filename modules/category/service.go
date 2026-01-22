package category

import "errors"

var categories []Category = []Category{
	{
		ID:          1,
		Name:        "Elektronik",
		Description: "Kategori barang-barang elektronik",
	},
	{
		ID:          2,
		Name:        "Fashion",
		Description: "Kategori barang-barang fashion",
	},
}

func validateCategory(category *Category) error {
	if category.Name == "" || category.Description == "" {
		return errors.New("name and description required")
	}

	return nil
}

func GetCategories() *[]Category {
	return &categories
}

func GetCategoryById(id int) *Category {
	for _, c := range categories {
		if c.ID == id {
			return &c
		}
	}

	return nil
}

func AddCategory(category *Category) (*Category, error) {
	if err := validateCategory(category); err != nil {
		return nil, err
	}

	category.ID = categories[len(categories)-1].ID + 1
	categories = append(categories, *category)

	return category, nil
}

func UpdateCategory(id int, category *Category) (*Category, error) {
	if err := validateCategory(category); err != nil {
		return nil, err
	}

	for i, c := range categories {
		if c.ID == id {
			category.ID = c.ID
			categories[i] = *category

			return category, nil
		}
	}

	return nil, errors.New("Category not found")
}

func DeleteCategory(id int) (*Category, error) {
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)

			return &c, nil
		}
	}

	return nil, errors.New("Category not found")
}
