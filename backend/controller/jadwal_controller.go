package controller

import (
	// "github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/service"
    "net/http"
	// "encoding/json"
	// "github.com/tapeds/fp-pbkk-golang/utils"
	// "github.com/google/uuid"
	// "time"

    "github.com/gin-gonic/gin"
	"fmt"
	// "log"
)

type JadwalController struct {
    jadwalService service.JadwalService
}

func NewJadwalController(service service.JadwalService) *JadwalController {
	return &JadwalController{
		jadwalService: service,
	}
}

// func (c *JadwalController) SearchPenerbangan(ctx *gin.Context) {
// 	// Get query parameters
// 	asal := ctx.DefaultQuery("asal", "")
// 	tujuan := ctx.DefaultQuery("tujuan", "")
// 	tanggalPerjalanan := ctx.DefaultQuery("tanggal_perjalanan", "")

// 	// Add your business logic here to query the database or services
// 	// For example, use the service to search for flights
// 	penerbanganList, err := c.jadwalService.GetAvailableFlights(asal, tujuan, tanggalPerjalanan)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
// 		return
// 	}

// 	// Return the response as JSON
// 	ctx.JSON(http.StatusOK, penerbanganList)
// }


func (c *JadwalController) SearchPenerbangan(ctx *gin.Context) {
	// Get query parameters
	// asal := ctx.DefaultQuery("asal", "")
	// tujuan := ctx.DefaultQuery("tujuan", "")
	tanggalPerjalanan := ctx.DefaultQuery("tanggal_perjalanan", "")
	// fmt.Println("Asal:", asal)
	// fmt.Println("Tujuan:", tujuan)
	fmt.Println("Tanggal Perjalanan:", tanggalPerjalanan)

	// Ensure the parameters are retrieved correctly
	// if asal == "" || tujuan == "" || tanggalPerjalanan == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameters"})
	// 	return
	// }
	if tanggalPerjalanan == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameters"})
		return
	}

	// // Parse the date (you may want to adjust the format)
	// tanggal, err := time.Parse("2006-01-02", tanggalPerjalanan)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
	// 	return
	// }

	// Use the service to query the database with the parameters
	penerbanganList, err := c.jadwalService.GetAvailableFlights(tanggalPerjalanan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
		return
	}

	fmt.Println("Parsed Tanggal:", tanggalPerjalanan)

	// Return the response as JSON
	ctx.JSON(http.StatusOK, penerbanganList)
}