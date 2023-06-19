package fiber

var BasicMiddleware = `package middleware

import "github.com/gofiber/fiber/v2"
	
func BasicMiddleware(c *fiber.Ctx) error {
	println("Middleware")
	err := c.Next()
	if err != nil {
		 panic(err)
	}
}`
