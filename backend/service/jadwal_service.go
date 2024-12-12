package service

import (
	// "fmt"

	// "github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/entity"
	"github.com/tapeds/fp-pbkk-golang/repository"

	// "context"
	// "errors"
	// "github.com/google/uuid"
	// "gorm.io/gorm"
)

type JadwalService interface {
	GetAvailableFlights(tanggalPerjalanan string) ([]entity.Penerbangan, error)

}

type jadwalService struct {
	ticketRepo    repository.TicketRepository
	passengerRepo repository.PassengerRepository
	penerbanganRepo repository.AdminRepository
	jadwalRepo repository.PenerbanganRepository
}

func NewJadwalService(jadwalRepo repository.PenerbanganRepository) JadwalService {
	return &jadwalService{
		jadwalRepo: jadwalRepo,
	}
}

func (s *jadwalService) GetAvailableFlights(tanggalPerjalanan string) ([]entity.Penerbangan, error) {
    return s.jadwalRepo.FindByQuery(tanggalPerjalanan)
}