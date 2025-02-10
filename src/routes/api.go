package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"lema/src/controllers"
	"lema/src/middlewares"
)

func Api(app *fiber.App) {

	app.Get("/", controllers.WelcomeHandler)
	route := app.Group("", logger.New())

	//user routes
	user := route.Group("users")
	user.Get("", middlewares.Paginate(), controllers.GetUsers)
	user.Get("/count", controllers.GetUserCount)
	user.Get("/:id", controllers.GetUserById)

	//post routes
	post := route.Group("posts")
	post.Post("", controllers.CreatePost)
	post.Get("", controllers.GetPosts)
	post.Delete("/:id", controllers.DeletePost)
}
