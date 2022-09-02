package storage

import (
	v1 "github.com/alinjie/go-by-example/internal/api/v1"
	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func dogToDogDTO(d Dog) v1.DogDTO {
	return v1.DogDTO{
		ID:   int(d.ID),
		Name: d.Name,
	}
}

func (s Storage) ListDogs() ([]v1.DogDTO, error) {
	out := []v1.DogDTO{}
	dogs := []Dog{}
	if err := s.db.Find(&dogs).Error; err != nil {
		return nil, err
	}

	for _, d := range dogs {
		out = append(out, dogToDogDTO(d))
	}

	return out, nil
}

func (s Storage) GetDog(id int) (v1.DogDTO, error) {
	dog := Dog{}

	if err := s.db.First(&dog, id).Error; err != nil {
		return v1.DogDTO{}, err
	}

	return dogToDogDTO(dog), nil
}
