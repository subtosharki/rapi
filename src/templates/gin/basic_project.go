package gin

func BasicProject(packageName string) string {
	return `package ` + packageName + `

import (
	"github.com/gin-gonic/gin"
	"` + packageName + `/src/middlewares"
	"` + packageName + `/src/routes"
)

func main() {
	r := gin.Default()
	r.Use(middleware.BasicMiddleware())
	r.GET("/", routes.BasicRoute)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}`
}
