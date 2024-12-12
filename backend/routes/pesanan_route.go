package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/controller"
	"github.com/tapeds/fp-pbkk-golang/middleware"
	"github.com/tapeds/fp-pbkk-golang/service"
)

func Pesanan(route *gin.Engine, pesananController controller.PesananController, jwtService service.JWTService) {
	routes := route.Group("/api/pesanan")
	{
		routes.GET("", middleware.Authenticate(jwtService), pesananController.GetAllPenerbanganByUserID)
		// routes.GET("", pesananController.GetAllPenerbanganByUserID)
		routes.GET("/:id", middleware.Authenticate(jwtService), pesananController.GetTicketByID)
		// routes.GET("/:id", pesananController.GetTicketByID)
		// route.GET("/pesanan/:id", middleware.Authenticate(jwtService), pesananController.GetTicketByID)
		// router.GET("/pesanan", middleware.Authenticate(jwtService), pesananController.GetAllTicketWithPagination)
	}
}
