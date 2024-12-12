package repository

import (
	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
	"context"
)

type PassengerRepository interface {
	AddPenumpang(ctx context.Context, penumpang entity.Penumpang) (entity.Penumpang, error)
}

type passengerRepository struct {
	connection *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) PassengerRepository {
	return &passengerRepository{
		connection: db,
	}
}

func (db *passengerRepository) AddPenumpang(ctx context.Context, penumpang entity.Penumpang) (entity.Penumpang, error) {
	uc := db.connection.Create(&penumpang)
	if uc.Error != nil {
		return entity.Penumpang{}, uc.Error
	}
	return penumpang, nil
}

