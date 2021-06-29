package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/testfiber/connect"
	"github.com/testfiber/models"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 20
	offset := (page - 1) * limit
	var total int64

	var products []models.Product

	connect.Database.Preload("Category").Preload("Picture").Offset(offset).Limit(limit).Find(&products)

	connect.Database.Model(&models.Product{}).Count(&total)

	last_page := math.Ceil(float64(int(total) / limit))

	return c.JSON(fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": last_page,
		},
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	connect.Database.Create(&product)

	return c.JSON(product)

}

func GetProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	connect.Database.Preload("Category").Preload("Picture").Find(&product)

	return c.JSON(product)
}

func UpdateProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	connect.Database.Model(&product).Updates(product)

	return c.JSON("Update successfully")
}

func DeleteProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	connect.Database.Delete(&product)

	return c.JSON("Delete successfully")
}
