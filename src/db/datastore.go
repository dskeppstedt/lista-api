package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Db interface {
	Connect()
	GetAppName()
}

type Mongodb struct {
	Url string
}

func NewMongoDb(url string) (m *Mongodb) {
	m = &Mongodb{Url: url}
	return
}

type Info struct {
	Name    string `bson:"name" json:"name"`
	Version int    `bson:"version" json:"version"`
}

func (this *Mongodb) Connect() {
	tries := 5

	var session *mgo.Session
	var err error
	for tries != 0 {
		session, err = mgo.Dial(this.Url)
		if err != nil {
			log.Println(err)
			log.Println("Will retry", tries, "time(s)...")
			tries -= 1
		} else {
			break
		}
	}

	var result Info
	c := session.DB("lista").C("info")
	c.Find(nil).One(&result)
	log.Println(result)
}

func (this *Mongodb) GetAppName() {
	//TODO
}
