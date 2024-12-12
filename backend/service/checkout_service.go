package service

import (
	"fmt"

	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/entity"
	"github.com/tapeds/fp-pbkk-golang/repository"

	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckoutService interface {
	CreateTiket(userID uuid.UUID, request dto.CheckoutRequest) (*entity.Tiket, error)
	// GetTicketDetails(penerbanganID string) (*entity.Penerbangan, error)
	// GetTicketDetails(c context.Context, req dto.PenerbanganResponse, penerbanganID string) (*entity.Penerbangan, error)
	GetPenerbanganDetail(ctx context.Context, id uuid.UUID) (entity.Penerbangan, error)
	// FindTicketByPenerbanganID(penerbanganID string) (*entity.Tiket, error)

}

type checkoutService struct {
	ticketRepo    repository.TicketRepository
	passengerRepo repository.PassengerRepository
	penerbanganRepo repository.AdminRepository
}

func NewCheckoutService(ticketRepo repository.TicketRepository, passengerRepo repository.PassengerRepository, penerbanganRepo repository.AdminRepository) CheckoutService {
	return &checkoutService{
		ticketRepo:    ticketRepo,
		passengerRepo: passengerRepo,
		penerbanganRepo: penerbanganRepo,
	}
}

// func (s *checkoutService) CheckoutTiket(ctx context.Context, req dto.CheckoutRequest, userID uuid.UUID) (*entity.Tiket, error) {
// 	// Validasi input
// 	if len(req.Penumpang) == 0 {
// 		return nil, errors.New("minimal harus ada satu penumpang")
// 	}

// 	// Buat entitas Tiket
// 	tiket := entity.Tiket{
// 		ID:            uuid.New(),
// 		PenerbanganID: req.PenerbanganID,
// 		UserID:        userID,
// 	}

// 	// Tambahkan data penumpang ke Tiket
// 	for _, penumpang := range req.Penumpang {
// 		tiket.Penumpang = append(tiket.Penumpang, entity.Penumpang{
// 			ID:   uuid.New(),
// 			Name: penumpang.Name,
// 			NIK:  penumpang.NIK,
// 		})
// 	}

// 	// Simpan ke database
// 	if err := s.ticketRepo.CreateTiket(ctx, &tiket); err != nil {
// 		return nil, err
// 	}

// 	return &tiket, nil
// }

func (s *checkoutService) CreateTiket(userID uuid.UUID, request dto.CheckoutRequest) (*entity.Tiket, error) {
	ctx := context.Background()

	ticket := entity.Tiket{
		UserID:       userID,
		PenerbanganID: request.PenerbanganID,
	}
	savedTiket, err := s.ticketRepo.CreateTiket(ctx, ticket)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to create ticket")
	}

	for _, passenger := range request.Penumpangs {
		p := entity.Penumpang{
			TiketID: 	savedTiket.ID,
			Name:     	passenger.Name,
			NIK:      	passenger.NIK,
		}
		_, err := s.passengerRepo.AddPenumpang(ctx, p)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed to create passenger")
		}
	}

	reloadedTiket, err := s.ticketRepo.GetTiketWithPenumpangs(ctx, savedTiket.ID)
	if err != nil {
		return nil, errors.New("failed to load ticket with passengers")
	}

	return &reloadedTiket, nil
}

// func (s *checkoutService) GetTicketDetails(penerbanganID string) (*entity.Tiket, error) {
// 	ticket, err := s.ticketRepo.FindTicketByPenerbanganID(penerbanganID)
// 	if err != nil {
// 		return nil, errors.New("ticket not found")
// 	}
// 	return ticket, nil
// }

func (as *checkoutService) GetPenerbanganDetail(ctx context.Context, id uuid.UUID) (entity.Penerbangan, error) {
	// Fetch penerbangan by ID
	penerbangan, err := as.penerbanganRepo.GetPenerbanganByID(ctx, nil, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Penerbangan{}, errors.New("penerbangan not found")
		}
		return entity.Penerbangan{}, err
	}
	return penerbangan, nil
}
