package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tapeds/fp-pbkk-golang/entity"
)

const (
	MESSAGE_FAILED_CREATE_BANDARA       = "failed create bandara"
	MESSAGE_FAILED_CREATE_MASKAPAI      = "failed create maskapai"
	MESSAGE_FAILED_CREATE_PENERBANGAN   = "failed create penerbangan"
	MESSAGE_FAILED_GET_LIST_PENERBANGAN = "failed get list penerbangan"
	MESSAGE_FAILED_GET_LIST_MASKAPAI    = "failed get list maskapai"
	MESSAGE_FAILED_GET_LIST_BANDARA     = "failed get list bandara"
	MESSAGE_FAILED_EDIT_PENERBANGAN     = "failed edit penerbangan"
	MESSAGE_FAILED_EDIT_MASKAPAI        = "failed edit maskapai"
	MESSAGE_FAILED_EDIT_BANDARA         = "failed edit bandara"

	MESSAGE_SUCCESS_CREATE_BANDARA       = "success create bandara"
	MESSAGE_SUCCESS_CREATE_MASKAPAI      = "success create maskapai"
	MESSAGE_SUCCESS_CREATE_PENERBANGAN   = "success create penerbangan"
	MESSAGE_SUCCESS_GET_LIST_PENERBANGAN = "success get list penerbangan"
	MESSAGE_SUCCESS_GET_LIST_MASKAPAI    = "success get list maskapai"
	MESSAGE_SUCCESS_GET_LIST_BANDARA     = "success get list bandara"
	MESSAGE_SUCCESS_EDIT_PENERBANGAN     = "success edit penerbangan"
	MESSAGE_SUCCESS_EDIT_MASKAPAI        = "success edit maskapai"
	MESSAGE_SUCCESS_EDIT_BANDARA         = "success edit bandara"
)

var (
	ErrCreateBandara            = errors.New("failed to create bandara")
	ErrCreateMaskapai           = errors.New("failed to create maskapai")
	ErrBandaraAlreadyExists     = errors.New("bandara already exist")
	ErrPriceBelowZero           = errors.New("harga is zero or negative")
	ErrCapacityBelowZero        = errors.New("capacity is zero or negative")
	ErrScheduleUnmatch          = errors.New("jadwal datang is earlier than jadwal berangkat")
	ErrMatchingAirport          = errors.New("source and destination airport is the same")
	ErrPenerbanganAlreadyExists = errors.New("penerbangan already exist")
	ErrPenerbanganNotFound      = errors.New("penerbangan not found")
	ErrMaskapaiNotFound         = errors.New("maskapai not found")
	ErrEditMaskapai             = errors.New("failed to edit maskapai")
	ErrEditPenerbangan          = errors.New("failed to edit penerbangan")
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

	PenerbanganCreateRequest struct {
		NoPenerbangan      string    `json:"no_penerbangan" binding:"required"`
		JadwalBerangkat    time.Time `json:"jadwal_berangkat" binding:"required"`
		JadwalDatang       time.Time `json:"jadwal_datang" binding:"required"`
		Harga              int       `json:"harga" binding:"required"`
		Kapasitas          int       `json:"kapasitas" binding:"required"`
		BandaraBerangkatID uuid.UUID `json:"bandara_berangkat" binding:"required"`
		BandaraDatangID    uuid.UUID `json:"bandara_datang" binding:"required"`
		MaskapaiID         uuid.UUID `json:"maskapai" binding:"required"`
	}

	BandaraResponse struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Kode string    `json:"kode"`
		Kota string    `json:"kota"`
	}

	BandaraArahResponse struct {
		ID   uuid.UUID       `json:"id"`
		Name string          `json:"name"`
		Kode string          `json:"kode"`
		Kota string          `json:"kota"`
		Arah entity.ArahEnum `json:"arah"`
	}

	MaskapaiResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Image string    `json:"image"`
	}

	PenerbanganResponse struct {
		ID              uuid.UUID             `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
		NoPenerbangan   string                `json:"no_penerbangan"`
		JadwalBerangkat time.Time             `json:"jadwal_berangkat"`
		JadwalDatang    time.Time             `json:"jadwal_datang"`
		Harga           int                   `json:"harga"`
		Kapasitas       int                   `json:"kapasitas"`
		Maskapai        MaskapaiResponse      `json:"maskapai"`
		Bandaras        []BandaraArahResponse `json:"bandaras"`
	}

	PenerbanganPaginationResponse struct {
		Data []PenerbanganResponse `json:"data"`
		PaginationResponse
	}

	GetAllPenerbanganRepositoryResponse struct {
		Penerbangans []entity.Penerbangan
		PaginationResponse
	}

	PenerbanganEditRequest struct {
		ID                 uuid.UUID `json:"id" binding:"required"`
		NoPenerbangan      string    `json:"no_penerbangan" binding:"required"`
		JadwalBerangkat    time.Time `json:"jadwal_berangkat" binding:"required"`
		JadwalDatang       time.Time `json:"jadwal_datang" binding:"required"`
		Harga              int       `json:"harga" binding:"required"`
		Kapasitas          int       `json:"kapasitas" binding:"required"`
		BandaraBerangkatID uuid.UUID `json:"bandara_berangkat" binding:"required"`
		BandaraDatangID    uuid.UUID `json:"bandara_datang" binding:"required"`
		MaskapaiID         uuid.UUID `json:"maskapai" binding:"required"`
	}

	MaskapaiEditRequest struct {
		ID    uuid.UUID `json:"id" binding:"required"`
		Name  string    `json:"name" binding:"required"`
		Image string    `json:"image" binding:"required"`
	}

	BandaraEditRequest struct {
		ID   uuid.UUID `json:"id" binding:"required"`
		Name string    `json:"name" binding:"required"`
		Kode string    `json:"kode" binding:"required"`
		Kota string    `json:"kota" binding:"required"`
	}
)
