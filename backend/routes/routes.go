package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/controller"
	// "net/http"
	// "github.com/tapeds/fp-pbkk-golang/middleware"
	// "github.com/tapeds/fp-pbkk-golang/service"
)

func Routes(route *gin.Engine, adminController controller.AdminController, penerbanganController controller.JadwalController) {
	routes := route.Group("/api")
	{
		routes.GET("/list-jadwal", adminController.GetPenerbangan)
		routes.GET("/list-bandara", adminController.GetBandara)
		routes.GET("/list-maskapai", adminController.GetMaskapai)
		routes.GET("/", penerbanganController.SearchPenerbangan)

	}}