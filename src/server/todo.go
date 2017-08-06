package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lista/api/models"
	"lista/api/util"
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

	var requestTodo models.Todo
	error = json.Unmarshal(body, &requestTodo)
	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(response, "Malformed request")
	}

	newTodo := models.NewTodo(requestTodo.Title)

	//find the user
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	err := DbStore.CreateTodo(user.Email, *newTodo)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(response, "Could not save todo")
	}

}
