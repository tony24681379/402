package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type CarSpace struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"carSpaceId"`
	Name      string        `bson:"name" json:"name"`
	Location  *GeoJson      `bson:"location" json:"location"`
	State     int           `bson:"state" json:"state"`
	UserID    bson.ObjectId `bson:"userId,omitempty" json:"userId,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

type CarPark struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"carParkId"`
	Name      string        `bson:"name" json:"name"`
	Location  *GeoJson      `bson:"location" json:"location"`
	CarSpace  *CarSpace     `bson:"carSpace" json:"carSpace"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}
