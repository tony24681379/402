package mongo

import (
	"github.com/golang/glog"
	"github.com/tony24681379/402/model"
)

type LocationDAO struct {
	Mongo
}

func (d *LocationDAO) All(query map[string]interface{}) ([]*model.Location, error) {
	return nil, nil
}

func (d *LocationDAO) NewLocation(l *model.Location) error {
	ds := d.MongoSession.Copy()
	defer ds.Close()
	c := ds.DB(d.MongoDBName).C(model.ModelLocation)

	err := c.Insert(l)
	if err != nil {
		glog.Error(err)
		return err
	}

	return nil
}
