package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c *Controller) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
