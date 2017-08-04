package server

import (
	"encoding/json"
	"io/ioutil"
	"lista/api/models"
	"lista/api/util"
	"log"
	"net/http"

	"fmt"
)

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

	var user models.User
	error = json.Unmarshal(body, &user)
	if error != nil {
		log.Println("Could not parse body")
		response.WriteHeader(400)
		fmt.Fprintln(response, "Malformed body")
		return
	}

	//check that the user exists
	if !DbStore.ExistUser(user.Email) {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(response, "User does not exist")
		return
	}

	//TODO: or if refreshtoken is present use that..
	if len(user.Refresh) > 0 {
		if !DbStore.CorrectRefreshToken(user) {
			response.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(response, "Wrong username/password combination")
			return
		}
	} else {
		//and the password is correct
		if !DbStore.CorrectUserPassword(user) {
			response.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(response, "Wrong username/password combination")
			return
		}
	}

	//FROM HERE THE USER IS AUTHORIZED

	//create token and send it!
	token, error := util.CreateToken(user.Email)
	if error != nil {
		log.Println("Signing error")
		log.Println(error)
		response.WriteHeader(500)
		fmt.Fprintln(response, "A token could not be issued")
		return
	}

	usr, err := DbStore.GetUser(user.Email)
	if err != nil {
		log.Println("Error get user", err)
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(response, "Could not get user")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	//send both tokens back to client
	ut := models.NewUserTokens(token, usr.Refresh)
	json.NewEncoder(response).Encode(ut)
}
