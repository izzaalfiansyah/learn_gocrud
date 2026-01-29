package category

import "errors"

type Service struct {
	repo *Repository
}

func validateCategory(category *Category) error {
	if category.Name == "" || category.Description == "" {
		return errors.New("name and description required")
	}

	return nil
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) GetCategories() (*[]Category, error) {
	return service.repo.GetCategories()
}

func (service *Service) GetCategoryByID(id *int) (*Category, error) {
	return service.repo.GetCategoryByID(id)
}

func (service *Service) AddCategory(category *Category) (*Category, error) {
	err := validateCategory(category)
	if err != nil {
		return nil, err
	}

	return service.repo.CreateCategory(category)
}

func (service *Service) UpdateCategory(id *int, category *Category) (*Category, error) {
	err := validateCategory(category)
	if err != nil {
		return nil, err
	}

	return service.repo.UpdateCategory(id, category)
}

func (service *Service) DeleteCategory(id *int) error {
	return service.repo.DeleteCategory(id)
}
