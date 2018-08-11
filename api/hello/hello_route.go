package hello

import "github.com/gin-gonic/gin"

func InitRoutes(g *gin.Engine, c *Controller) {
	g.GET("/api/hello", c.Hello)
}
