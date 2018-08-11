package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tony24681379/402/api/hello"
)

func InitRoutes(g *gin.Engine) error {
	hello.InitRoutes(g, &hello.Controller{})

	return nil
}
