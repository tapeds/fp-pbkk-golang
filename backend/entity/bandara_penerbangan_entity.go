package entity

import (
	"github.com/google/uuid"
)

type ArahEnum string

const (
	ArahBerangkat ArahEnum = "BERANGKAT"
	ArahDatang    ArahEnum = "DATANG"
)

type BandaraPenerbangan struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BandaraID     uuid.UUID
	PenerbanganID uuid.UUID
	Arah          ArahEnum `gorm:"type:arah_enum"`

	Bandara Bandara

	Timestamp
}
