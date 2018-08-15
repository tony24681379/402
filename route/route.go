package route

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/golang/glog"
	"github.com/tony24681379/402/api/hello"
	"github.com/tony24681379/402/api/location"
	"github.com/tony24681379/402/mongo"
)

func InitRoutes(g *gin.Engine) error {
	mongoURL := "127.0.0.1"
	mongoDBName := "402"

	s, err := mgo.Dial(mongoURL)
	if err != nil {
		glog.Error(err)
		return err
	}

	hello.InitRoutes(g, &hello.Controller{})
	location.InitRoutes(g, &location.Controller{
		LocationDAO: &mongo.LocationDAO{
			Mongo: mongo.Mongo{
				MongoSession: s,
				MongoDBName:  mongoDBName,
			},
		},
	})

	return nil
}
