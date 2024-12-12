package controller

import (
	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/service"
    "net/http"
	"github.com/tapeds/fp-pbkk-golang/utils"
	"github.com/google/uuid"

    "github.com/gin-gonic/gin"
	"log"
)

type CheckoutController struct {
    CheckoutService service.CheckoutService
}

func NewCheckoutController(service service.CheckoutService) *CheckoutController {
	return &CheckoutController{
		CheckoutService: service,
	}
}

func (cc *CheckoutController) CreateTicket(c *gin.Context) {
	// Get user_id from the context (JWT Token)
	userIdInterface, exists := c.Get("user_id")
	if !exists {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "user_id not found in context", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// log.Printf("user_id from context: %v, type: %T", userIdInterface, userIdInterface)

	userIdString, ok := userIdInterface.(string)
	if !ok {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "user_id is not a valid string", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	
	userID, err := uuid.Parse(userIdString)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "invalid user_id format", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	penerbanganIDParam := c.Param("id")
	penerbanganID, err := uuid.Parse(penerbanganIDParam)
	if err != nil {
		log.Printf("Invalid penerbanganID: %v", err)
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "invalid penerbangan_id format", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	// log.Printf("userID from context: %v, type: %T", userID, userID)
	// Bind the request body to a DTO
	var request dto.CheckoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	request.PenerbanganID = penerbanganID

	if err := utils.ValidateStruct(request); err != nil {
		response := utils.BuildResponseFailed("Validation failed", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call service to create the ticket and passengers
	ticket, err := cc.CheckoutService.CreateTiket(userID, request)
	if err != nil {
		response := utils.BuildResponseFailed("Failed to create ticket", err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Return success response
	response := utils.BuildResponseSuccess("Ticket and passengers created successfully", ticket)
	c.JSON(http.StatusOK, response)
}

func (cc *CheckoutController) ShowCheckoutForm(c *gin.Context) {
	// Get penerbanganID from URL params
	// Parse the ID from URL parameter
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	// Get penerbangan details from service
	penerbangan, err := cc.CheckoutService.GetPenerbanganDetail(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the penerbangan details
	c.JSON(http.StatusOK, gin.H{"data": penerbangan})
}
