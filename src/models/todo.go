package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Title string        `json:"title"`
	Done  bool          `json:"done"`
}

func NewTodo(title string) *Todo {
	return &Todo{bson.NewObjectId(), title, false}
}
