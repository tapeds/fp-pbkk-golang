package service

import (
	// "fmt"

	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/entity"
	"github.com/tapeds/fp-pbkk-golang/repository"

	"context"
	"errors"
	"github.com/google/uuid"
)

type PesananService interface {
	GetTicketByID(ctx context.Context, id uuid.UUID)(entity.Tiket, error)
	GetAllTicketWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TicketPaginationResponse, error)
	GetTicketDetails(penerbanganID string) (*entity.Tiket, error)
	GetPenerbanganByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Penerbangan, error)
}

type pesananService struct {
	ticketRepo    	repository.TicketRepository
	jwtService 		JWTService
	penerbanganRepo 	repository.PenerbanganRepository
	// passengerRepo repository.PassengerRepository
}

func NewPesananService(ticketRepo repository.TicketRepository, jwtService JWTService, penerbanganRepo repository.PenerbanganRepository) PesananService {
	return &pesananService{
		ticketRepo:    ticketRepo,
		jwtService: jwtService,
		penerbanganRepo: penerbanganRepo,
	}
}

// func (s *pesananService) GetTicketByID(ctx context.Context, tiketID string) (dto.CheckoutResponse, error) {
// 	tiket, err := s.ticketRepo.GetTicketByID(tiketID)
// 	if err != nil {
// 		return dto.CheckoutResponse{}, errors.New("ticket not found")
// 	}

// 	ticketResponse := dto.CheckoutResponse{
// 		ID:             tiket.ID,
// 		PenerbanganID:  tiket.PenerbanganID,
// 	}

// 	return ticketResponse, nil
// }

func (s *pesananService) GetTicketByID(ctx context.Context, id uuid.UUID) (entity.Tiket, error) {
	tiket, err := s.ticketRepo.GetTicketByID(ctx, nil, id)
	if err != nil {
		return entity.Tiket{}, errors.New("ticket not found")
	}
	return tiket, nil
}

func (s *pesananService) GetAllTicketWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TicketPaginationResponse, error) {
	dataWithPaginate, err := s.ticketRepo.GetAllTicketWithPagination(ctx, nil, req)
	if err != nil {
		return dto.TicketPaginationResponse{}, err
	}

	var datas []dto.CheckoutResponse
	for _, tiket := range dataWithPaginate.Tickets {
		var penumpangs []dto.Penumpang
		for _, p := range tiket.Penumpang {
			penumpang := dto.Penumpang{
				Name: p.Name,
				NIK:  p.NIK,
			}
			penumpangs = append(penumpangs, penumpang)
		}
		data := dto.CheckoutResponse{
			ID : 				tiket.ID,
			PenerbanganID:  	tiket.PenerbanganID,
			Penumpangs:     	penumpangs,  
		}

		datas = append(datas, data)
	}

	return dto.TicketPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}


func (s *pesananService) GetTicketDetails(tiketID string) (*entity.Tiket, error) {
	// Memanggil repository untuk mengambil tiket berdasarkan penerbanganID
	ticket, err := s.ticketRepo.FindTicketByID(tiketID)
	if err != nil {
		return nil, errors.New("ticket not found")
	}
	// Mengembalikan tiket dengan data penumpang yang terkait
	return ticket, nil
}

func (s *pesananService) GetPenerbanganByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Penerbangan, error) {
    return s.penerbanganRepo.FindByUserID(ctx, userID)
}

