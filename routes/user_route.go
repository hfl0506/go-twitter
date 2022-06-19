package routes

import (
	"go-twitter/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router) {
	userRoute := api.Group("/users")

	userRoute.Get("/", controllers.GetUsers)
	userRoute.Get("/:id", controllers.GetUser)
	userRoute.Post("/", controllers.CreateUser)
	userRoute.Put("/:id", controllers.UpdateUser)
	userRoute.Delete("/:id", controllers.DeleteUser)
}
