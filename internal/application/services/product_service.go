package services

import (
	"github.com/profadevairvitorio/go_with_DDD/internal/domain/entities"
	"github.com/profadevairvitorio/go_with_DDD/internal/domain/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *entities.Product) error {
	return s.repo.Save(product)
}

func (s *ProductService) GetAllProducts() ([]*entities.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) FindProductByID(id int) (*entities.Product, error) {
	return s.repo.FindByID(id)
}
