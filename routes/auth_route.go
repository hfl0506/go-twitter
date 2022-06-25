package routes

import (
	"go-twitter/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	authRoute := api.Group("/auth")

	authRoute.Post("/login", controllers.Login)
	//authRoute.Post("/forgetpassword")
	//authRoute.Post("/activate")
}
