package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/testfiber/connect"
	"github.com/testfiber/models"
)

func AllCategories(c *fiber.Ctx) error {
	var categories []models.Category

	connect.Database.Find(&categories)

	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return err
	}

	connect.Database.Create(&category)

	return c.JSON(category)

}

func GetCategoryById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}

	connect.Database.Find(&category)

	return c.JSON(category)
}

func UpdateCategoryById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}

	if err := c.BodyParser(&category); err != nil {
		return err
	}

	connect.Database.Model(&category).Updates(category)

	return c.JSON("Update successfully")
}

func DeleteCategoryById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category := models.Category{
		Id: uint(id),
	}

	connect.Database.Delete(&category)

	return c.JSON("Delete successfully")
}
