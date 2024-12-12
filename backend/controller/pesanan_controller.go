package controller

import (
	// "context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/service"
	"github.com/tapeds/fp-pbkk-golang/utils"
	"github.com/google/uuid"
	"context"
)

type (
	PesananController interface {
		GetAllTicket(ctx *gin.Context)
		GetTicketByID(ctx *gin.Context)
		GetAllPenerbanganByUserID(ctx *gin.Context)
	}

	pesananController struct {
		pesananService service.PesananService
	}
)

func NewPesananController(us service.PesananService) PesananController {
	return &pesananController{
		pesananService: us,
	}
}

func (c *pesananController) GetTicketByID(ctx *gin.Context) {
	tiketID := ctx.Param("id")
	// id, err := uuid.Parse(idParam)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket ID"})
	// 	return
	// }

	ticket, err := c.pesananService.GetTicketDetails(tiketID)
	if err != nil {
		response := utils.BuildResponseFailed("Failed to fetch ticket details", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.BuildResponseSuccess("Ticket details fetched successfully", ticket)
	ctx.JSON(http.StatusOK, response)
}

func (c *pesananController) GetAllTicket(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.pesananService.GetAllTicketWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

// func (c *pesananController) GetAllPenerbanganByUserID(ctx *gin.Context) {
//     tokenString := ctx.GetHeader("Authorization")
//     userID, err := utils.DecodeToken(tokenString) // Decode JWT to get the user ID
//     if err != nil {
//         ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//         return
//     }

//     penerbangans, err := c.pesananService.GetPenerbanganByUserID(context.Background(), userID)
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     ctx.JSON(http.StatusOK, penerbangans)
// }

// func (c *pesananController) GetAllPenerbanganByUserID(ctx *gin.Context) {
//     userIdInterface, exists := ctx.Get("user_id")
//     if !exists {
//         response := utils.BuildResponseFailed("Error", "User ID not found in context", nil)
//         ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
//         return
//     }

// 	userIdString, ok := userIdInterface.(string)
// 	if !ok {
// 		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "user_id is not a valid string", nil)
// 		ctx.JSON(http.StatusUnauthorized, response)
// 		return
// 	}

// 	userID, err := uuid.Parse(userIdString)
// 	if err != nil {
// 		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "invalid user_id format", nil)
// 		ctx.JSON(http.StatusUnauthorized, response)
// 		return
// 	}

//     penerbangans, err := c.pesananService.GetPenerbanganByUserID(context.Background(), userID.(uuid.UUID))
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     ctx.JSON(http.StatusOK, penerbangans)
// }

func (c *pesananController) GetAllPenerbanganByUserID(ctx *gin.Context) {
    // Ambil user_id dari context
    userIdInterface, exists := ctx.Get("user_id")
    if !exists {
        response := utils.BuildResponseFailed("Error", "User ID not found in context", nil)
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
        return
    }

    // Pastikan user_id adalah string
    userIdString, ok := userIdInterface.(string)
    if !ok {
        response := utils.BuildResponseFailed("Error", "User ID is not a valid string", nil)
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
        return
    }

    // Konversi string ke uuid.UUID
    userID, err := uuid.Parse(userIdString)
    if err != nil {
        response := utils.BuildResponseFailed("Error", "Invalid user ID format", nil)
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
        return
    }

    // Panggil service untuk mendapatkan penerbangan berdasarkan user_id
    penerbangans, err := c.pesananService.GetPenerbanganByUserID(context.Background(), userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Kirimkan response dengan data penerbangan
    ctx.JSON(http.StatusOK, gin.H{"data": penerbangans})
}
