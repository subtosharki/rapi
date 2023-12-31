package fiber

func BasicProject(packageName string) string {
	return `package main

import (
	"github.com/gofiber/fiber/v2"
	"` + packageName + `/src/middlewares"
	"` + packageName + `/src/routes"
)

var PORT = ":3000"


func main() {
	app := fiber.New()

	app.Get("/", routes.BasicRoute)
	app.Use("/", middlewares.BasicMiddleware)

	err := app.Listen("0.0.0.0" + PORT)
	if err != nil {
		panic(err)
	}
}`
}
