package gin

func BasicMiddleware(middlewareName string, middlewarePathName string) string {
	return `package ` + middlewarePathName + `

import "github.com/gin-gonic/gin"

func ` + middlewareName + `() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Middleware")
		c.Next()
	}
}`
}
