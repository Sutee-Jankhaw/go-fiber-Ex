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
	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "23012023",
		},
	})
	user.Get("/", c.GetUser)
	user.Get("/json", auth, c.GetUserJson)
	user.Get("/search", auth, c.SearchUser)
	user.Get("/:id", auth, c.GetUserById)
	user.Post("/", auth, c.CreateUser)
	user.Put("/:id", auth, c.UpdateUser)
	user.Delete("/:id", auth, c.RemoveUser)
}
