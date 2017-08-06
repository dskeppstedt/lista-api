package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Title string
	Done  bool
}

func NewTodo(title string) *Todo {
	return &Todo{bson.NewObjectId(), title, false}
}
