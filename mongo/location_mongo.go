package mongo

import (
	"time"

	"github.com/globalsign/mgo"
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

func (d *LocationDAO) NewLocation(l *model.Location) (*model.Location, error) {
	ds := d.MongoSession.Copy()
	defer ds.Close()
	c := ds.DB(d.MongoDBName).C(model.ModelLocation)

	l.ID = bson.NewObjectId()
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()

	change := mgo.Change{
		Update:    l,
		Upsert:    true,
		ReturnNew: true,
	}
	_, err := c.Find(bson.M{"_id": l.ID}).Apply(change, &l)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return l, nil
}

func (d *LocationDAO) GetLocation(long, lat float64, state, scope int) ([]*model.Location, error) {
	ds := d.MongoSession.Copy()
	defer ds.Close()
	c := ds.DB(d.MongoDBName).C(model.ModelLocation)
	location := []*model.Location{}
	match := bson.M{}
	if state != 2 {
		match["state"] = state
	}
	match["geo"] = bson.M{
		"$nearSphere": bson.M{
			"$geometry": bson.M{
				"type":        "Point",
				"coordinates": []float64{long, lat},
			},
			"$maxDistance": scope,
		},
	}
	glog.Info(match)
	err := c.Find(match).All(&location)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return location, nil
}
