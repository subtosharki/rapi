package echo

func BasicMiddleware(routeName string, routePathName string) string {
	return `package ` + routePathName + `

import "github.com/gofiber/fiber/v2"

func ` + routeName + `() {
	func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
}`
}
