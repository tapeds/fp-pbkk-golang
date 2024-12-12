package routes

import (
	"github.com/tapeds/fp-pbkk-golang/controller"
	"github.com/tapeds/fp-pbkk-golang/service"
    "github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/middleware"
)

func Checkout(router *gin.Engine, checkoutController *controller.CheckoutController, jwtService service.JWTService) {
	routes := router.Group("/api/checkout")
	{
		routes.POST("/:id", middleware.Authenticate(jwtService), checkoutController.CreateTicket)
		routes.GET("/:id", middleware.Authenticate(jwtService), checkoutController.ShowCheckoutForm)
		// routes.POST("/:id", checkoutController.CreateTicket)
		// routes.GET("/:id", checkoutController.ShowCheckoutForm)
	}
}
