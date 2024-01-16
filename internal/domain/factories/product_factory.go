package factories

import (
	"errors"
	"github.com/profadevairvitorio/go_with_DDD/internal/domain/entities"
)

func NewProduct(name string, price float64) (*entities.Product, error) {
	if name == "" || price <= 0 {
		return nil, errors.New("invalid product details")
	}
	return &entities.Product{Name: name, Price: price}, nil
}
