package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

const ModelUsers string = "users"

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"userId"`
	Name      string        `bson:"name" json:"name"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}
