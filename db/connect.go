package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	MongoDBUrl = "mongodb://192.168.137.140:27017/xtest"
)

func Connect() {
	mongo, err := mgo.ParseURL(MongoDBUrl)
	s, err := mgo.Dial(MongoDBUrl)
	if err != nil {
		log.Printf("[info] Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	s.SetSafe(&mgo.Safe{})
	log.Println("[info] Connected to", MongoDBUrl)
	Session = s
	Mongo = mongo
}
