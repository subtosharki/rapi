package fiber

func BasicMiddleware(routeName string, routePathName string) string {
	return `package ` + routePathName + `

import "github.com/gofiber/fiber/v2"
	
func ` + routeName + `(c *fiber.Ctx) {
	println("Middleware")
	err := c.Next()
	if err != nil {
		 panic(err)
	}
}`
}
