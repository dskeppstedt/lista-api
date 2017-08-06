package db

import (
	"fmt"
	"lista/api/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Db interface {
	Connect()
	GetAppInfo()
	CreateUser()
	GetUser()
	UpdateUserWithTokens()
	ExistUser()
	CorrectUserPassword()
	CorrectRefreshToken()
	//Todos
	CreateTodo(string, models.Todo)
	ReadTodos(string)
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

func (this *Mongodb) CreateUser(user models.User) error {
	session := this.Session
	collection := session.DB("lista").C("users")
	err := collection.Insert(user)
	return err
}

func (this *Mongodb) GetUser(email string) (models.User, error) {
	session := this.Session
	c := session.DB("lista").C("users")
	result := models.User{}
	err := c.Find(bson.M{"email": email}).One(&result)
	return result, err
}

func (this *Mongodb) ExistUser(email string) bool {
	session := this.Session
	c := session.DB("lista").C("users")

	result := models.User{}
	err := c.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println("Found user", result)
	return true
}

//TODO: refactor cuz code smells
func (this *Mongodb) CorrectUserPassword(user models.User) bool {
	session := this.Session
	c := session.DB("lista").C("users")
	result := models.User{}
	err := c.Find(bson.M{"email": user.Email}).One(&result)
	if err != nil {
		log.Println("CorrectUserPassword", err)
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return false
	}

	return true
}

func (this *Mongodb) UpdateUserWithToken(user models.User, refresh string) bool {
	session := this.Session
	c := session.DB("lista").C("users")

	userQuery := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{"refresh_token": refresh}}

	err := c.Update(userQuery, update)
	if err != nil {
		log.Println("Update could not be performed", err)
		return false
	}
	return true
}

func (this *Mongodb) CorrectRefreshToken(user models.User) bool {
	session := this.Session
	c := session.DB("lista").C("users")
	result := models.User{}
	err := c.Find(bson.M{"email": user.Email}).One(&result)
	if err != nil {
		log.Println("Could not find user", err)
		return false
	}

	return result.Refresh == user.Refresh
}

//TODOS

func (this *Mongodb) CreateTodo(user string, todo models.Todo) error {
	session := this.Session
	c := session.DB("lista").C("users")
	userQuery := bson.M{"email": user}
	updateQuery := bson.M{"$push": bson.M{"todos": todo}}
	err := c.Update(userQuery, updateQuery)
	return err
}

func (this *Mongodb) ReadTodos(email string) ([]models.Todo, error) {
	user, err := this.GetUser(email)
	return user.Todos, err
}

func (this *Mongodb) DeleteTodo(email string, id string) error {
	s := this.Session
	c := s.DB("lista").C("users")
	userQuery := bson.M{"email": email}
	updateQuery := bson.M{"$pull": bson.M{"todos": bson.M{"_id": bson.ObjectIdHex(id)}}}
	err := c.Update(userQuery, updateQuery)
	log.Println("heler", err)
	return err
}

func (this *Mongodb) UpdateTodo(email string, id string, todo models.Todo) error {
	s := this.Session
	c := s.DB("lista").C("users")
	userQuery := bson.M{"email": email, "todos._id": bson.ObjectIdHex(id)}
	updateQuery := bson.M{"$set": bson.M{"todos.$.title": todo.Title, "todos.$.done": todo.Done}}
	err := c.Update(userQuery, updateQuery)
	return err
}
