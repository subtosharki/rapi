package gin

var BasicRoute = `package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicRoute(c *gin.Context) error {
	c.String(http.StatusOK, "Hello, World!")
}`
