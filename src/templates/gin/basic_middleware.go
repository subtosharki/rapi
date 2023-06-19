package gin

var BasicMiddleware = `package middleware

import "github.com/gin-gonic/gin"

func BasicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Middleware")
		c.Next()
	}
}`
