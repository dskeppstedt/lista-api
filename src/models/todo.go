package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string        `json:"title,omitempty"`
	Done  bool          `json:"done,omitempty"`
}

func NewTodo(title string) *Todo {
	return &Todo{bson.NewObjectId(), title, false}
}
