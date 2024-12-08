package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tapeds/fp-pbkk-golang/entity"
)

const (
	MESSAGE_FAILED_CREATE_BANDARA  = "failed create bandara"
	MESSAGE_FAILED_CREATE_MASKAPAI = "failed create maskapai"

	MESSAGE_SUCCESS_CREATE_BANDARA  = "success create bandara"
	MESSAGE_SUCCESS_CREATE_MASKAPAI = "success create maskapai"
)

var (
	ErrCreateBandara        = errors.New("failed to create bandara")
	ErrCreateMaskapai       = errors.New("failed to create maskapai")
	ErrBandaraAlreadyExists = errors.New("bandara already exist")
)

type (
	BandaraCreateRequest struct {
		Name string `json:"name" binding:"required"`
		Kode string `json:"kode" binding:"required"`
		Kota string `json:"kota" binding:"required"`
	}

	MaskapaiCreateRequest struct {
		Name  string `json:"name" binding:"required"`
		Image string `json:"image" binding:"required"`
	}

	BandaraResponse struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Kode string    `json:"kode"`
		Kota string    `json:"kota"`
	}

	MaskapaiResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Image string    `json:"image"`
	}

	PenerbanganResponse struct {
		ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		NoPenerbangan    string    `json:"no_penerbangan"`
		JadwalBerangkat  time.Time `json:"jadwal_berangkat"`
		JadwalDatang     time.Time `json:"jadwal_datang"`
		Harga            int       `json:"harga"`
		Kapasitas        int       `json:"kapasitas"`
		Maskapai         string    `json:"maskapai"`
		BandaraBerangkat string    `json:"bandara_berangkat"`
		BandaraDatang    string    `json:"bandara_datang"`
	}

	PenerbanganPaginationResponse struct {
		Data []PenerbanganResponse `json:"data"`
		PaginationResponse
	}

	GetAllPenerbanganRepositoryResponse struct {
		Penerbangans []entity.Penerbangan
		PaginationResponse
	}
)
