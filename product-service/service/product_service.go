package service

import (
	"product-service/models"
	"product-service/repository"
)

type ProductService interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	UpdateProduct(id string, product models.Product) error
	DeleteProduct(id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	return s.repo.Create(product)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetProductByID(id string) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateProduct(id string, product models.Product) error {
	return s.repo.Update(id, product)
}

func (s *productService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}