package main

import (
	"log"
	"os"

	"github.com/tapeds/fp-pbkk-golang/command"
	"github.com/tapeds/fp-pbkk-golang/config"
	"github.com/tapeds/fp-pbkk-golang/controller"
	"github.com/tapeds/fp-pbkk-golang/middleware"
	"github.com/tapeds/fp-pbkk-golang/repository"
	"github.com/tapeds/fp-pbkk-golang/routes"
	"github.com/tapeds/fp-pbkk-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		flag := command.Commands(db)
		if !flag {
			return
		}
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository  repository.UserRepository  = repository.NewUserRepository(db)
		adminRepository repository.AdminRepository = repository.NewAdminRepository(db)

		// Service
		userService  service.UserService  = service.NewUserService(userRepository, jwtService)
		adminService service.AdminService = service.NewAdminService(adminRepository, jwtService)

		// Controller
		userController  controller.UserController  = controller.NewUserController(userService)
		adminController controller.AdminController = controller.NewAdminController(adminService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	routes.Admin(server, adminController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
