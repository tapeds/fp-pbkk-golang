package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/controller"
	"github.com/tapeds/fp-pbkk-golang/middleware"
	"github.com/tapeds/fp-pbkk-golang/service"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		// User
		routes.POST("/register", userController.Register)
		routes.GET("/getAll", userController.GetAllUser)
		routes.POST("/login", userController.Login)
		routes.DELETE("", middleware.Authenticate(jwtService), userController.Delete)
		routes.PATCH("", middleware.Authenticate(jwtService), userController.Update)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		routes.POST("/verify_email", userController.VerifyEmail)
		routes.POST("/send_verification_email", userController.SendVerificationEmail)
	}
}
