package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("user")

	user.Get("/", c.GetUser)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	}))

	user.Post("/", c.CreateUser)
	user.Put("/:id", c.UpdateUser)
	user.Delete("/:id", c.RemoveUser)
}
