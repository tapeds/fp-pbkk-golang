package entity

import (
	"time"

	"github.com/google/uuid"
)

type Penerbangan struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NoPenerbangan   string    `json:"no_penerbangan"`
	JadwalBerangkat time.Time `json:"jadwal_berangkat"`
	JadwalDatang    time.Time `json:"jadwal_datang"`
	Harga           int       `json:"harga"`
	Kapasitas       int       `json:"kapasitas"`

	MaskapaiID         uuid.UUID
	Maskapai           Maskapai
	Tiket              []Tiket
	BandaraPenerbangan []BandaraPenerbangan

	Timestamp
}
