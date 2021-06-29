package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/testfiber/connect"
	"github.com/testfiber/models"
)

func AllCarts(c *fiber.Ctx) error {
	var carts []models.Cart

	connect.Database.Preload("Product").Find(&carts)

	return c.JSON(carts)
}

func CreateCart(c *fiber.Ctx) error {
	var cart models.Cart

	if err := c.BodyParser(&cart); err != nil {
		return err
	}

	connect.Database.Create(&cart)

	return c.JSON(cart)

}

func GetCartById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	cart := models.Cart{
		Id: uint(id),
	}

	connect.Database.Preload("Product").Find(&cart)

	return c.JSON(cart)
}

func UpdateCartById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	cart := models.Cart{
		Id: uint(id),
	}

	if err := c.BodyParser(&cart); err != nil {
		return err
	}

	connect.Database.Model(&cart).Updates(cart)

	return c.JSON("Update successfully")
}

func DeleteCartById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	cart := models.Cart{
		Id: uint(id),
	}

	connect.Database.Delete(&cart)

	return c.JSON("Delete successfully")
}
