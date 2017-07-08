package models

//Info represents information about the webapp
//stored in the database
type Info struct {
	Name    string `bson:"name" json:"name"`
	Version int    `bson:"version" json:"version"`
}
