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

	l.Geo.Coordinates = []float64{l.Geo.Long, l.Geo.Lat}
	l, err := c.LocationDAO.NewLocation(l)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, l)
}

func (c *Controller) GetLocation(ctx *gin.Context) {
	distance, err := strconv.Atoi(ctx.DefaultQuery("distance", "5000"))

	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	state, err := strconv.Atoi(ctx.DefaultQuery("state", "2"))
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

	locations, err := c.LocationDAO.GetLocation(long, lat, state, distance)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, loc := range locations {
		loc.Geo.Long = loc.Geo.Coordinates[0]
		loc.Geo.Lat = loc.Geo.Coordinates[1]
		loc.Geo.Coordinates = []float64{}
	}
	ctx.JSON(http.StatusOK, locations)
}

func (c *Controller) SetState(ctx *gin.Context) {

	l := &model.Location{}
	if err := ctx.Bind(l); err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	l.Geo.Coordinates = []float64{l.Geo.Long, l.Geo.Lat}
	err := c.LocationDAO.UpdateLocation(l)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "OK")

}
