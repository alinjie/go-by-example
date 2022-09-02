package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) Storage {
	return Storage{db: db}
}

func (s Storage) Migrate() error {
	return s.db.AutoMigrate(&Dog{})
}

func (s Storage) Seed() error {
	dogs := []Dog{
		{Name: "Mira"},
		{Name: "Lassie"},
		{Name: "Nikolai"},
	}

	if err := s.db.Create(&dogs).Error; err != nil {
		return err
	}

	return nil
}
