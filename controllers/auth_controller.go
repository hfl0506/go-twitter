package controllers

import (
	"go-twitter/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(ctx *fiber.Ctx) error {
	var req *LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user, _ := GetUserByName(req.Name)

	isMatch, err := utils.ComparePassword(req.Password, user.Password)

	if !isMatch && err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err})
	}

	arg := utils.Payload{
		ID:   string(user.ID),
		Name: user.Name,
	}

	accessToken, _ := utils.GenerateAccessToken(&arg)

	return ctx.JSON(fiber.Map{"accessToken": accessToken})
}
