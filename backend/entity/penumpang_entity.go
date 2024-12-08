package entity

import (
	"github.com/google/uuid"
)

type Penumpang struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name    string    `json:"name"`
	NIK     string    `json:"nik"`
	TiketID uuid.UUID

	Timestamp
}
