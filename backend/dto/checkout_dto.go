package dto

import (
	// "errors"

	"github.com/google/uuid"
	"github.com/tapeds/fp-pbkk-golang/entity"
)

// CheckoutRequest represents the request structure for creating a ticket
type CheckoutRequest struct {
	ID			 	uuid.UUID    `json:"id"` // ID of the ticket
	PenerbanganID  	uuid.UUID    `json:"penerbangan_id" validate:"required"`
	Penumpangs     	[]Penumpang  `json:"penumpang" binding:"required"`
}

// Passenger represents the passenger details in the request
type Penumpang struct {
	// ID   uuid.UUID `json:"id"`
	Name string `json:"name" binding:"required"`
	NIK  string `json:"nik" binding:"required"`
}

type CheckoutResponse struct {
    ID			 	uuid.UUID    `json:"id"` // ID of the ticket
	PenerbanganID  	uuid.UUID    `json:"penerbangan_id" binding:"required"`
	Penumpangs     	[]Penumpang  `json:"penumpang" binding:"required"`
}

type GetAllTicketRepositoryResponse struct {
	Tickets []entity.Tiket
	PaginationResponse
}

type TicketPaginationResponse struct {
	Data []CheckoutResponse `json:"data"`
	PaginationResponse
}
