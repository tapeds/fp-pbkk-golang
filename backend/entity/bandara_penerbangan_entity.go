package entity

import (
	"github.com/google/uuid"
)

type BandaraPenerbangan struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BandaraID     uuid.UUID
	PenerbanganID uuid.UUID

	Timestamp
}
