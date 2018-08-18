package location

import "github.com/gin-gonic/gin"

func InitRoutes(g *gin.Engine, c *Controller) {
	g.POST("/api/location", c.NewLocation)
	g.GET("/api/location", c.GetLocation)
	g.PATCH("/api/location", c.SetState)
}
