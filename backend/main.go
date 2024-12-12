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
		ticketRepository repository.TicketRepository = repository.NewTicketRepository(db)
		passengerRepository repository.PassengerRepository = repository.NewPassengerRepository(db)
		penerbanganRepository repository.PenerbanganRepository = repository.NewPenerbanganRepository(db)

		// Service
		userService  service.UserService  = service.NewUserService(userRepository, jwtService)
		adminService service.AdminService = service.NewAdminService(adminRepository, jwtService)
		checkoutService service.CheckoutService = service.NewCheckoutService(ticketRepository, passengerRepository, adminRepository)
		pesananService service.PesananService = service.NewPesananService(ticketRepository, jwtService, penerbanganRepository)
		jadwalService service.JadwalService = service.NewJadwalService(penerbanganRepository)

		// Controller
		userController  controller.UserController  = controller.NewUserController(userService)
		adminController controller.AdminController = controller.NewAdminController(adminService)
		checkoutController *controller.CheckoutController = controller.NewCheckoutController(checkoutService)
		pesananController controller.PesananController = controller.NewPesananController(pesananService)
		jadwalController controller.JadwalController = *controller.NewJadwalController(jadwalService)
		// userRepository repository.UserRepository = repository.NewUserRepository(db)

		// Service
		// userService service.UserService = service.NewUserService(userRepository, jwtService)

		// Controller
		// userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	routes.Admin(server, adminController, jwtService)
	routes.Checkout(server, checkoutController, jwtService)
	routes.Pesanan(server, pesananController, jwtService)
	routes.Routes(server, adminController, jadwalController)

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
