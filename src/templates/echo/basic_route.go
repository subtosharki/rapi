package echo

func BasicRoute(routeName string, packageName string) string {
	return `package ` + packageName + `

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ` + routeName + `(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}`
}
