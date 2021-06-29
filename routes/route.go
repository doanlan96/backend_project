package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/testfiber/controllers"
	"github.com/testfiber/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/:name", controllers.SayName)
	app.Get("/api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUserById)
	app.Put("/api/users/:id", controllers.UpdateUserById)
	app.Delete("/api/users/:id", controllers.DeleteUserById)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRoleById)
	app.Put("/api/roles/:id", controllers.UpdateRoleById)
	app.Delete("/api/roles/:id", controllers.DeleteRoleById)

	app.Get("/api/products", controllers.AllProducts)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProductById)
	app.Put("/api/products/:id", controllers.UpdateProductById)
	app.Delete("/api/products/:id", controllers.DeleteProductById)

	app.Get("/api/categories", controllers.AllCategories)
	app.Post("/api/categories", controllers.CreateCategory)
	app.Get("/api/categories/:id", controllers.GetCategoryById)
	app.Put("/api/categories/:id", controllers.UpdateCategoryById)
	app.Delete("/api/categories/:id", controllers.DeleteCategoryById)

	app.Get("/api/pictures", controllers.AllPictures)
	app.Post("/api/pictures", controllers.CreatePicture)
	app.Get("/api/pictures/:id", controllers.GetPictureById)
	app.Put("/api/pictures/:id", controllers.UpdatePictureById)
	app.Delete("/api/pictures/:id", controllers.DeletePictureById)

	app.Get("/api/cart", controllers.AllCarts)
	app.Post("/api/cart", controllers.CreateCart)
	app.Get("/api/cart/:id", controllers.GetCartById)
	app.Put("/api/cart/:id", controllers.UpdateCartById)
	app.Delete("/api/cart/:id", controllers.DeleteCartById)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/upload", "./upload")
}
