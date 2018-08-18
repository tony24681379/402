package location

import (
	"net/http"
	"strconv"
	"strings"

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

func (c *Controller) GetLocation(ctx *gin.Context) {
	distance, err := strconv.Atoi(ctx.Query("distance"))
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	geo := ctx.Query("geo")
	geoS := strings.Split(geo, ",")
	long, err := strconv.ParseFloat(geoS[0], 64)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lat, err := strconv.ParseFloat(geoS[1], 64)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location, err := c.LocationDAO.GetLocation(long, lat, distance)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, location)
}
