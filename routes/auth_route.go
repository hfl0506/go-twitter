package routes

import "github.com/gofiber/fiber/v2"

func SetupAuthRoutes(api fiber.Router) {
	authRoute := api.Group("/auth")

	authRoute.Post("/login")
	authRoute.Post("/forgetpassword")
	authRoute.Post("/activate")
}
