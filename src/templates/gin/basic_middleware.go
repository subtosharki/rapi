package gin

func BasicMiddleware(routeName string, routePathName string) string {
	return `package ` + routePathName + `

import "github.com/gin-gonic/gin"

func ` + routeName + `() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Middleware")
		c.Next()
	}
}`
}
