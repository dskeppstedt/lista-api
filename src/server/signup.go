package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lista/api/models"
	"lista/api/util"
	"log"
	"net/http"
)

func signup(response http.ResponseWriter, request *http.Request) {

	body, error := ioutil.ReadAll(request.Body)
	if error != nil {
		log.Println("Could not read body")
		response.WriteHeader(500)
		fmt.Fprintln(response, "Can not read the request body")
		return
	}

	var user models.User
	error = json.Unmarshal(body, &user)
	if error != nil {
		log.Println("Could not parse body")
		response.WriteHeader(400)
		fmt.Fprintln(response, "Malformed body")
		return
	}

	//okay two things can happen now, either this is a new user or
	//someone tries to sign up using a existing email.

	//check if that user already exists
	if DbStore.ExistUser(user.Email) {
		response.WriteHeader(http.StatusConflict)
		fmt.Fprintln(response, "User with that email already exisit")
		return
	}
	//encrypt password
	user.Password = util.GenerateHash(user.Password)
	log.Println(user)
	//create the user

	if err := DbStore.CreateUser(user); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(response, "Could not save user")
		return
	}

	//else.. create jwt token and refresh token and send that along

	//create refresh token

	//store refresh token with the user

	//create jwt token

	//send both tokens back to client

	ut := models.NewUserTokens("a", "b")

	json.NewEncoder(response).Encode(ut)

}
