package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 15
	offset := (page - 1) * limit
	//var total int64

	//var products []Product
	//var products interface{}

	//db.Offset(offset).Limit(limit).Find(&products)

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)
	db.Model(&Product{}).Count(&total)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}
}
