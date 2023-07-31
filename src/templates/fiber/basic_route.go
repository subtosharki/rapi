package fiber

func BasicRoute(routeName string, routePathName string) string {
	return `package ` + routePathName + `

import "github.com/gofiber/fiber/v2"
	
func ` + routeName + `(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}`
}
