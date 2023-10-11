package controllers

import (
	"awesomeProject/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
	}

	user.LastName = "Doe344"

	// return c.SendString("Hello, World 👋!")
	return c.JSON(user)
}
