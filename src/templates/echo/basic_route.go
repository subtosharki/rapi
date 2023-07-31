package echo

func BasicRoute(routeName string, packageName string) string {
	return `package ` + packageName + `

import "github.com/labstack/echo/v4"

func ` + routeName + `(c echo.Context) error {
	println("request to /users")
	return next(c)
}`
}
