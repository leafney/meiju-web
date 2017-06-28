package models

import (
	""
)

const(
	CollectionVideo="videos"
)

type Video struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title" binding:"required"`
	

}
