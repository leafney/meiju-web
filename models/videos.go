package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionVideo = "videos"
)

type Video struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Title string        `json:"title" bson:"title" binding:"required"`
}
