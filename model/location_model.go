package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

const ModelLocation string = "locations"

type Location struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Geo       *Geo          `bson:"geo" json:"geo"`
	CarSpace  `bson:",inline" json:",inline"`
	CarPark   `bson:",inline" json:",inline"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type CarSpace struct {
	Name   string        `bson:"name,omitempty" json:"name,omitempty"`
	State  int           `bson:"state,omitempty" json:"state"`
	UserID bson.ObjectId `bson:"userId,omitempty" json:"userId,omitempty"`
}

type CarPark struct {
	CarSpace []*CarSpace `bson:"carSpace,omitempty" json:"carSpace,omitempty"`
}

type Geo struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
