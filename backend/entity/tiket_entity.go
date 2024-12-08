package entity

import (
	"github.com/google/uuid"
)

type Tiket struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID        uuid.UUID
	PenerbanganID uuid.UUID
	Penumpang     []Penumpang

	Timestamp
}
