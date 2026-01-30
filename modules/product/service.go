package product

import "errors"

type Service struct {
	repo *Repository
}

func validateProduct(category *Product) error {
	if category.Name == "" {
		return errors.New("name required")
	}

	return nil
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) GetProducts() (*[]Product, error) {
	return service.repo.GetProducts()
}

func (service *Service) GetProductByID(id *int) (*Product, error) {
	return service.repo.GetProductByID(id)
}

func (service *Service) AddProduct(category *Product) (*Product, error) {
	err := validateProduct(category)
	if err != nil {
		return nil, err
	}

	return service.repo.CreateProduct(category)
}

func (service *Service) UpdateProduct(id *int, category *Product) (*Product, error) {
	err := validateProduct(category)
	if err != nil {
		return nil, err
	}

	return service.repo.UpdateProduct(id, category)
}

func (service *Service) DeleteProduct(id *int) error {
	return service.repo.DeleteProduct(id)
}
