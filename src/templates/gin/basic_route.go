package gin

func BasicRoute(routeName string, routePathName string) string {
	return `package ` + routePathName + `

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ` + routeName + `(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
`
}
