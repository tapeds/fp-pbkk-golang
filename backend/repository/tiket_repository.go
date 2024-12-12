package repository

import (
	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
	"github.com/tapeds/fp-pbkk-golang/dto"
	"context"
	"math"
	"errors"
	"github.com/google/uuid"
)

type TicketRepository interface {
	CreateTiket(ctx context.Context, tiket entity.Tiket) (entity.Tiket, error)
	GetTicketByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Tiket, error)
	// GetAllTickets(ctx context.Context) ([]entity.Tiket, error)
	GetAllTicketWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllTicketRepositoryResponse, error)
	GetTiketWithPenumpangs(ctx context.Context, tiketID uuid.UUID) (entity.Tiket, error)
	// FindTicketByPenerbanganID(penerbanganID string) (*entity.Tiket, error)
	FindTicketByPenerbanganID(penerbanganID string) (*entity.Tiket, error)
	FindTicketByID(tiketID string) (*entity.Tiket, error)
	FindPenerbanganByUserID(ctx context.Context, userID string) ([]entity.Tiket, error)
}

type ticketRepository struct {
	connection *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		connection: db,
	}
}


func (db *ticketRepository) CreateTiket(ctx context.Context, tiket entity.Tiket) (entity.Tiket, error) {
	uc := db.connection.Create(&tiket)
	if uc.Error != nil {
		return entity.Tiket{}, uc.Error
	}
	// Load passengers after creating the ticket
	err := db.connection.Preload("Penumpang").First(&tiket, "id = ?", tiket.ID).Error
	if err != nil {
		return entity.Tiket{}, err
	}

	return tiket, nil
}

func (db *ticketRepository) GetTicketByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Tiket, error) {
	if tx == nil {
		tx = db.connection
	}

	var tiket entity.Tiket
	if err := tx.WithContext(ctx).Where("id = ?", id).Take(&tiket).Error; err != nil {
		return entity.Tiket{}, err
	}

	return tiket, nil
}

func (db *ticketRepository) GetAllTicketWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllTicketRepositoryResponse, error) {
	if tx == nil {
		tx = db.connection
	}

	var tickets []entity.Tiket
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Tiket{}).Count(&count).Error; err != nil {
		return dto.GetAllTicketRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&tickets).Error; err != nil {
		return dto.GetAllTicketRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllTicketRepositoryResponse{
		Tickets:     tickets,
		PaginationResponse: dto.PaginationResponse{
			Page: 		 req.Page,
			PerPage: 	 req.PerPage,
			Count: 		 count,
			MaxPage: 	 totalPage,
		},
	}, err
}

func (db *ticketRepository) GetTiketWithPenumpangs(ctx context.Context, tiketID uuid.UUID) (entity.Tiket, error) {
	var tiket entity.Tiket
	err := db.connection.Preload("Penumpang").First(&tiket, "id = ?", tiketID).Error
	if err != nil {
		return entity.Tiket{}, err
	}
	return tiket, nil
}

func (db *ticketRepository) FindTicketByPenerbanganID(penerbanganID string) (*entity.Tiket, error) {
	var ticket entity.Tiket
	if err := db.connection.Where("penerbangan_id = ?", penerbanganID).Preload("Penumpang").First(&ticket).Error; err != nil {
		return nil, errors.New("ticket not found")
	}
	return &ticket, nil
}

func (db *ticketRepository) FindTicketByID(tiketID string) (*entity.Tiket, error) {
	var ticket entity.Tiket
	if err := db.connection.Where("id = ?", tiketID).Preload("Penumpang").First(&ticket).Error; err != nil {
		return nil, errors.New("ticket not found")
	}
	return &ticket, nil
}

func (db *ticketRepository) FindPenerbanganByUserID(ctx context.Context, userID string) ([]entity.Tiket, error) {
    var tiket []entity.Tiket

    err := db.connection.WithContext(ctx).
        Where("user_id = ?", userID).
        Preload("Penerbangan.Maskapai").
        Preload("Penerbangan.BandaraPenerbangan.Bandara").
        Preload("Penumpang").
        Find(&tiket).Error

    if err != nil {
        return nil, err
    }
	db.connection = db.connection.Debug() // Tambahkan ini untuk mencetak query SQL


    return tiket, nil
}
