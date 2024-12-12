package repository

import (
    "context"
	// "math"

	"github.com/google/uuid"
	// "github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/entity"
	"gorm.io/gorm"
)

type PenerbanganRepository interface {
    FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Penerbangan, error)
    FindByQuery(tanggalPerjalanan string) ([]entity.Penerbangan, error)
}

type penerbanganRepository struct {
    db *gorm.DB
}

func NewPenerbanganRepository(db *gorm.DB) PenerbanganRepository {
    return &penerbanganRepository{db: db}
}

func (r *penerbanganRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Penerbangan, error) {
    var penerbangans []entity.Penerbangan
    // Mengambil data penerbangan berdasarkan UserID yang terdapat pada Tiket
    err := r.db.WithContext(ctx).
        Preload("Tiket").
        Preload("Penerbangan").
        Preload("Maskapai").
        Joins("JOIN tikets ON tikets.penerbangan_id = penerbangans.id").
        Where("tikets.user_id = ?", userID).
        Find(&penerbangans).Error

    if err != nil {
        return nil, err
    }
    return penerbangans, nil
}

func (r *penerbanganRepository) FindByQuery(tanggalPerjalanan string) ([]entity.Penerbangan, error) {
    var penerbangans []entity.Penerbangan
    query := r.db.Where("jadwal_berangkat like ?", tanggalPerjalanan)
    result := query.Find(&penerbangans)
    return penerbangans, result.Error
}
