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

func CreateTodo(response http.ResponseWriter, request *http.Request) {

	//read request body
	body, error := ioutil.ReadAll(request.Body)
	if error != nil {

		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(response, "Could not read request")
		return
	}

	var newTodo models.Todo
	error = json.Unmarshal(body, &newTodo)
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(response, "Malformed request")
	}

	//find the user
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	err := DbStore.CreateTodo(user.Email, newTodo)
	log.Println(err)
}
