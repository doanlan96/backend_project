package controllers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func SayName(c *fiber.Ctx) error {
	name, err := url.PathUnescape(c.Params("name"))
	fmt.Println(err)
	fmt.Println(name)

	msg := fmt.Sprintf("Hello, %s !", name)
	return c.SendString(msg)
}
