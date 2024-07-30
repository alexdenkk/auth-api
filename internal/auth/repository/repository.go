package repository

import "gorm.io/gorm"

// Repository - repository layer struct
type Repository struct {
	DB *gorm.DB
}

// New - function for creating Repository instance
func New(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
