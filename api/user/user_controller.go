package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tony24681379/402/model"
	"github.com/tony24681379/402/mongo"
)

type Controller struct {
	UserDAO *mongo.UserDAO
}

func (c *Controller) NewUser(ctx *gin.Context) {
	u := &model.User{}

	if err := ctx.Bind(u); err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.UserDAO.NewUser(u)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.String(http.StatusOK, "OK")
}
func (c *Controller) Login(ctx *gin.Context) {
	name := ctx.Param("name")
	user, err := c.UserDAO.FindUser(name)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user != nil {
		ctx.String(http.StatusOK, "OK")

	}
	ctx.String(http.StatusBadRequest, "")

}
