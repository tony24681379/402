package user

import "github.com/gin-gonic/gin"

func InitRoutes(g *gin.Engine, c *Controller) {
	g.POST("/api/user", c.NewUser)
	g.GET("/api/user/:name", c.Login)
}
