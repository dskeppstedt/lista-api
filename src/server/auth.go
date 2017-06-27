package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"time"

	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Email    string
	Password string
}

func auth(response http.ResponseWriter, request *http.Request) {
	//make sure that this is a post request
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(response, "Only POST is allowed")
		return
	}

	//read and parse the body, can probly extracted into middleware..
	body, error := ioutil.ReadAll(request.Body)
	if error != nil {
		log.Println("Could not read body")
		response.WriteHeader(500)
		fmt.Fprintln(response, "Can not read the request body")
		return
	}

	log.Println(string(body))

	var user User
	error = json.Unmarshal(body, &user)
	if error != nil {
		log.Println("Could not parse body")
		response.WriteHeader(400)
		fmt.Fprintln(response, "Malformed body")
		return
	}

	//check that the user exists and the password is correct

	if user.Email != "email@email.com" || user.Password != "123" {
		response.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(response, "User not authorized")
		return
	}

	//create token and send it!
	token, error := createToken(user.Email)
	if error != nil {
		log.Println("Signing error")
		log.Println(error)
		response.WriteHeader(500)
		fmt.Fprintln(response, "A token could not be issued")
		return
	}

	response.Header().Set("Content-Type", "application/jwt")
	response.WriteHeader(200)
	fmt.Fprintln(response, token)
}

//source https://github.com/dgrijalva/jwt-go/blob/master/http_example_test.go

type UserClaims struct {
	*jwt.StandardClaims
	Email string
}

func createToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &UserClaims{
		&jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 10).Unix()},
		email}
	return token.SignedString([]byte("foobar"))
}
