package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/testfiber/connect"
	"github.com/testfiber/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	connect.Database.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	connect.Database.Create(&role)

	return c.JSON(role)

}

func GetRoleById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	connect.Database.Find(&role)

	return c.JSON(role)
}

func UpdateRoleById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	connect.Database.Model(&role).Updates(role)

	return c.JSON("Update successfully")
}

func DeleteRoleById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	connect.Database.Delete(&role)

	return c.JSON("Delete successfully")
}
