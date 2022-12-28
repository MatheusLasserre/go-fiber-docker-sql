package routes

import (
	"github.com/MatheusLasserre/go-fiber-docker-sqloback/handlers"
	"github.com/MatheusLasserre/go-fiber-docker-sqloback/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Index)

	courses := app.Group("/courses", middleware.CoursesM)

	courses.Get("/", handlers.GetCourses)
	courses.Post("/", handlers.PostCourses)
}
