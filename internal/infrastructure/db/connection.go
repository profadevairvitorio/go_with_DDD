package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", "foo_bar")
}
