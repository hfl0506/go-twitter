package controllers

import (
	"go-twitter/entities"
	"go-twitter/services"
	"go-twitter/utils"

	"github.com/gofiber/fiber/v2"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required, alphanum"`
	Password string `json:"password" binding:"required, min=6"`
	Email    string `json:"email" binding:"required, email"`
}

func CreateUser(ctx *fiber.Ctx) error {
	var req *createUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	password, _ := utils.HashPassword(req.Password)

	arg := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
	}

	user, err := services.CreateUser(&arg)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(user)
}

func GetUsers(ctx *fiber.Ctx) error {
	users, err := services.GetUsers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(users)
}

func GetUser(ctx *fiber.Ctx) error {
	user, err := services.GetUser(ctx.Params("id"))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user *entities.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := services.UpdateUser(user, ctx.Params("id"))

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	user, err := services.DeleteUser(ctx.Params("id"))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(user)
}
