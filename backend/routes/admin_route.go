package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/controller"
	"github.com/tapeds/fp-pbkk-golang/middleware"
	"github.com/tapeds/fp-pbkk-golang/service"
)

func Admin(route *gin.Engine, adminController controller.AdminController, jwtService service.JWTService) {
	routes := route.Group("/api/admin")
	{
		routes.GET("/penerbangan", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.GetPenerbangan)
		routes.GET("/bandara", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.GetBandara)
		routes.GET("/maskapai", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.GetMaskapai)
		routes.POST("/bandara", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.AddBandara)
		routes.POST("/maskapai", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.AddMaskapai)
		routes.POST("/penerbangan", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.AddPenerbangan)
		routes.PATCH("/penerbangan", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.EditPenerbangan)
		routes.PATCH("/bandara", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.EditBandara)
		routes.PATCH("/maskapai", middleware.Authenticate(jwtService, middleware.WithRole("admin")), adminController.EditMaskapai)
	}
}
