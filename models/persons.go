package models

/*
	mongodb db:xtest  collection:xtest model person{Name,Phone}

*/

const (
	CollectionPerson = "xtest"
)

type Person struct {
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
}
