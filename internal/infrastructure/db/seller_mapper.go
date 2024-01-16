package db

import "github.com/profadevairvitorio/go_with_DDD/internal/domain/entities"

// ToDBSeller maps domain Seller entity to DB persistence model.
func ToDBSeller(seller *entities.Seller) *Seller {
	s := &Seller{
		Name: seller.Name,
	}
	s.ID = seller.ID

	return s
}

// FromDBSeller maps DB persistence model to domain Seller entity.
func FromDBSeller(dbSeller *Seller) *entities.Seller {
	s := &entities.Seller{
		ID:   dbSeller.ID,
		Name: dbSeller.Name,
	}
	s.ID = dbSeller.ID

	return s
}
