package controllers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
