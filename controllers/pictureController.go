package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/testfiber/connect"
	"github.com/testfiber/models"
)

func AllPictures(c *fiber.Ctx) error {
	var pictures []models.Picture

	connect.Database.Find(&pictures)

	return c.JSON(pictures)
}

func CreatePicture(c *fiber.Ctx) error {
	var picture models.Picture

	if err := c.BodyParser(&picture); err != nil {
		return err
	}

	connect.Database.Create(&picture)

	return c.JSON(picture)

}

func GetPictureById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	picture := models.Picture{
		Id: uint(id),
	}

	connect.Database.Find(&picture)

	return c.JSON(picture)
}

func UpdatePictureById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	picture := models.Picture{
		Id: uint(id),
	}

	if err := c.BodyParser(&picture); err != nil {
		return err
	}

	connect.Database.Model(&picture).Updates(picture)

	return c.JSON("Update successfully")
}

func DeletePictureById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	picture := models.Picture{
		Id: uint(id),
	}

	connect.Database.Delete(&picture)

	return c.JSON("Delete successfully")
}
