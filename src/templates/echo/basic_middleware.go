package echo

func BasicMiddleware(middlewareName string, middlewarePathName string) string {
	return `package ` + middlewarePathName + `

import "github.com/labstack/echo/v4"

func ` + middlewareName + `() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("Middleware")
			return next(c)
		}
	}
}`
}
