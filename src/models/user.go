package models

type User struct {
	Email    string
	Password string
	Refresh  string `bson:"refresh_token" json:"refresh"`
	Todos    []Todo `bson:"todos"`
}
