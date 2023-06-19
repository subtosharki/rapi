package fiber

var BasicRoute = `package routes

import "github.com/gofiber/fiber/v2"

func BasicRoute(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}`
