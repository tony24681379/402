package mongo

import (
	"github.com/globalsign/mgo/bson"
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

func (d *LocationDAO) GetLocation(long, lat float64, scope int) ([]*model.Location, error) {
	ds := d.MongoSession.Copy()
	defer ds.Close()
	c := ds.DB(d.MongoDBName).C(model.ModelLocation)

	location := []*model.Location{}
	err := c.Find(bson.M{
		"geo": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"$maxDistance": scope,
			},
		},
	}).All(&location)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return location, nil
}
