package factories

import (
	"errors"
	"github.com/profadevairvitorio/go_with_DDD/internal/domain/entities"
)

func NewSeller(name string) (*entities.Seller, error) {
	if name == "" {
		return nil, errors.New("invalid seller details")
	}
	return &entities.Seller{Name: name}, nil
}
