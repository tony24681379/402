package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tony24681379/402/model"
	"github.com/tony24681379/402/mongo"
)

type Controller struct {
	LocationDAO *mongo.LocationDAO
}

func (c *Controller) NewLocation(ctx *gin.Context) {
	l := &model.Location{}

	if err := ctx.Bind(l); err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.LocationDAO.NewLocation(l)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "OK")
}
