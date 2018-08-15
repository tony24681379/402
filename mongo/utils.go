package mongo

import "github.com/globalsign/mgo"

type Mongo struct {
	MongoSession *mgo.Session
	MongoDBName  string
}
