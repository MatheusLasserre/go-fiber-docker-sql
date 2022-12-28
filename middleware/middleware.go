package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func CoursesM(c *fiber.Ctx) error {
	log.Println("Here's the Courses Middleware!")
	return c.Next()
}
