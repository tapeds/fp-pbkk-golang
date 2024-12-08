package entity

import (
	"github.com/google/uuid"
)

type Bandara struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name string    `json:"name"`
	Kode string    `json:"kode"`
	Kota string    `json:"kota"`

	BandaraPenerbangan []BandaraPenerbangan

	Timestamp
}
