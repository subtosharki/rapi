package echo

func BasicProject(packageName string) string {
	return `package main

import (
  "github.com/labstack/echo/v4"
  "` + packageName + `/src/middlewares"
  "` + packageName + `/src/routes"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middlewares.BasicMiddleware())

  // Routes
  e.GET("/", routes.BasicRoute)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}
`
}
