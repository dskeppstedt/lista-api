package db

import (
	"lista/api/models"
	"log"

	"gopkg.in/mgo.v2"
)

type Db interface {
	Connect()
	GetAppName()
}

type Mongodb struct {
	Url     string
	Session *mgo.Session
}

func NewMongoDb(url string) (m *Mongodb) {
	m = &Mongodb{Url: url}
	return
}

//Connect establish a connection to mongodb.
//If it fails, it will retry five times.
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
	this.Session = session
}

func (this *Mongodb) GetAppInfo() models.Info {
	session := this.Session
	var result models.Info
	c := session.DB("lista").C("info")
	c.Find(nil).One(&result)
	log.Println(result)
	return result
}
