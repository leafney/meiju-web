package db

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	MongoDBUrl = "mongodb://localhost:27017/xtest"
)

func Connect() {
	mongo, err := mgo.ParseURL(MongoDBUrl)
	s, err := mgo.Dial(MongoDBUrl)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	s.SetSafe(&mgo.Safe{})
	log.Println("Connected to", MongoDBUrl)
	Session = s
	Mongo = mongo
}
