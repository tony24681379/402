package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/golang/glog"
	"github.com/tony24681379/402/model"
)

type UserDAO struct {
	Mongo
}

func (d *UserDAO) NewUser(u *model.User) error {
	ds := d.MongoSession.Copy()
	defer ds.Close()
	c := ds.DB(d.MongoDBName).C(model.ModelUsers)

	err := c.Insert(u)
	if err != nil {
		glog.Error(err)
		return err
	}

	return nil
}
func (d *UserDAO) FindUser(name string) (*model.User, error) {
	ds := d.MongoSession.Copy()
	u := &model.User{}
	defer ds.Close()
	c := ds.DB(d.MongoDBName).C(model.ModelUsers)

	err := c.Find(bson.M{"_id": name}).One(&u)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return u, nil
}
