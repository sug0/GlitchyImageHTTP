package database

import (
	"log"

	"github.com/Toyz/GlitchyImageHTTP/core"
	"github.com/globalsign/mgo"
)

type mongo struct {
	mgoSession *mgo.Session
	database   string
	collection string
}

var MongoInstance *mongo

func NewMongo() {
	session, err := mgo.Dial(core.GetEnv("MONGO_HOST", "localhost"))
	if err != nil {
		log.Println(err)
	}
	session.SetMode(mgo.Monotonic, true)

	MongoInstance = &mongo{
		mgoSession: session,
		database:   core.GetEnv("MONGO_DB", "glitch"),
		collection: core.GetEnv("MONGO_COLLECTION", "artIds"),
	}
}

func (mg *mongo) GetSession() *mgo.Session {
	return mg.mgoSession.Copy()
}

func (mg *mongo) GetCollection() (*mgo.Session, *mgo.Collection) {
	session := mg.GetSession()
	c := session.DB(mg.database).C(mg.collection)

	return session, c
}
