package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/controller"
	"github.com/tapeds/fp-pbkk-golang/service"
)

func Admin(route *gin.Engine, adminController controller.AdminController, jwtService service.JWTService) {
	routes := route.Group("/api/admin")
	{
		routes.GET("/penerbangan", adminController.GetPenerbangan)
		routes.POST("/bandara", adminController.AddBandara)
		routes.POST("/maskapai", adminController.AddMaskapai)
	}
}
