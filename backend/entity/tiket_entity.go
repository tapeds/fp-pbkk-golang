package entity

import (
	"github.com/google/uuid"
)

type Tiket struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	PenerbanganID uuid.UUID `json:"penerbangan_id"`
	
	Penumpang     []Penumpang `gorm:"foreignKey:TiketID" json:"penumpang"`
	UserID        uuid.UUID 
	User          User
	Timestamp
}
