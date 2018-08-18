package route

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/golang/glog"
	"github.com/tony24681379/402/api/hello"
	"github.com/tony24681379/402/api/location"
	"github.com/tony24681379/402/api/user"
	"github.com/tony24681379/402/model"

	"github.com/tony24681379/402/mongo"
)

func InitRoutes(g *gin.Engine, mongoURL string) error {
	mongoDBName := "402"

	s, err := mgo.Dial(mongoURL)
	if err != nil {
		glog.Error(err)
		return err
	}
	index := mgo.Index{
		Key: []string{"$2dsphere:geo"},
	}
	if err := s.DB(mongoDBName).C(model.ModelLocation).EnsureIndex(index); err != nil {
		glog.Fatal(err)
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
	user.InitRoutes(g, &user.Controller{
		UserDAO: &mongo.UserDAO{
			Mongo: mongo.Mongo{
				MongoSession: s,
				MongoDBName:  mongoDBName,
			},
		},
	})

	return nil
}
