package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/testfiber/connect"
	"github.com/testfiber/routes"
)

func main() {
	connect.Connect()

	app := fiber.New() //trả về 1 biến kiểu *fiber.App

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
