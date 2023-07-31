package fiber

func BasicMiddleware(middlewareName string, middlewarePathName string) string {
	return `package ` + middlewarePathName + `

import "github.com/gofiber/fiber/v2"
	
func ` + middlewareName + `(c *fiber.Ctx) {
	println("Middleware")
	err := c.Next()
	if err != nil {
		 panic(err)
	}
}`
}
