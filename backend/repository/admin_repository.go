package repository

import (
	"context"
	"math"

	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
)

type (
	AdminRepository interface {
		GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPenerbanganRepositoryResponse, error)
		GetAllBandara(ctx context.Context, tx *gorm.DB) ([]entity.Bandara, error)
		GetAllMaskapai(ctx context.Context, tx *gorm.DB) ([]entity.Maskapai, error)
		CheckBandaraCode(ctx context.Context, tx *gorm.DB, kode string) (entity.Bandara, bool, error)
		CheckPenerbanganNumber(ctx context.Context, tx *gorm.DB, number string) (entity.Penerbangan, bool, error)
		CreateBandara(ctx context.Context, tx *gorm.DB, bandara entity.Bandara) (entity.Bandara, error)
		CreateMaskapai(ctx context.Context, tx *gorm.DB, maskapai entity.Maskapai) (entity.Maskapai, error)
		CreatePenerbangan(ctx context.Context, tx *gorm.DB, penerbangan entity.Penerbangan) (entity.Penerbangan, error)
	}

	adminRepository struct {
		db *gorm.DB
	}
)

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (ar *adminRepository) GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllPenerbanganRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var penerbangans []entity.Penerbangan
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Penerbangan{}).Count(&count).Error; err != nil {
		return dto.GetAllPenerbanganRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Preload("BandaraPenerbangan.Bandara").Preload("Maskapai").Scopes(Paginate(req.Page, req.PerPage)).Find(&penerbangans).Error; err != nil {
		return dto.GetAllPenerbanganRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllPenerbanganRepositoryResponse{
		Penerbangans: penerbangans,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (ar *adminRepository) CheckBandaraCode(ctx context.Context, tx *gorm.DB, kode string) (entity.Bandara, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var bandara entity.Bandara
	if err := tx.WithContext(ctx).Where("kode = ?", kode).Take(&bandara).Error; err != nil {
		return entity.Bandara{}, false, err
	}

	return bandara, true, nil
}

func (ar *adminRepository) GetAllBandara(ctx context.Context, tx *gorm.DB) ([]entity.Bandara, error) {
	if tx == nil {
		tx = ar.db
	}

	var bandara []entity.Bandara
	if err := tx.WithContext(ctx).Find(&bandara).Error; err != nil {
		return []entity.Bandara{}, err
	}

	return bandara, nil
}

func (ar *adminRepository) GetAllMaskapai(ctx context.Context, tx *gorm.DB) ([]entity.Maskapai, error) {
	if tx == nil {
		tx = ar.db
	}

	var maskapai []entity.Maskapai
	if err := tx.WithContext(ctx).Find(&maskapai).Error; err != nil {
		return []entity.Maskapai{}, err
	}

	return maskapai, nil
}

func (ar *adminRepository) CheckPenerbanganNumber(ctx context.Context, tx *gorm.DB, number string) (entity.Penerbangan, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var penerbangan entity.Penerbangan
	if err := tx.WithContext(ctx).Where("no_penerbangan = ?", number).Take(&penerbangan).Error; err != nil {
		return entity.Penerbangan{}, false, err
	}

	return penerbangan, true, nil
}

func (ar *adminRepository) CreateBandara(ctx context.Context, tx *gorm.DB, bandara entity.Bandara) (entity.Bandara, error) {
	if tx == nil {
		tx = ar.db
	}

	if err := tx.WithContext(ctx).Create(&bandara).Error; err != nil {
		return entity.Bandara{}, err
	}

	return bandara, nil
}

func (ar *adminRepository) CreateMaskapai(ctx context.Context, tx *gorm.DB, maskapai entity.Maskapai) (entity.Maskapai, error) {
	if tx == nil {
		tx = ar.db
	}

	if err := tx.WithContext(ctx).Create(&maskapai).Error; err != nil {
		return entity.Maskapai{}, err
	}

	return maskapai, nil
}

func (ar *adminRepository) CreatePenerbangan(ctx context.Context, tx *gorm.DB, penerbangan entity.Penerbangan) (entity.Penerbangan, error) {
	if tx == nil {
		tx = ar.db
	}

	if err := tx.WithContext(ctx).Create(&penerbangan).Error; err != nil {
		return entity.Penerbangan{}, err
	}

	if err := tx.WithContext(ctx).
		Preload("Maskapai").
		Preload("BandaraPenerbangan.Bandara").
		First(&penerbangan, penerbangan.ID).Error; err != nil {
		return entity.Penerbangan{}, err
	}

	return penerbangan, nil
}
