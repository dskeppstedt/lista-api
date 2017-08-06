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
		return
	}

	newTodo := models.NewTodo(requestTodo.Title)

	//find the user
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	err := DbStore.CreateTodo(user.Email, *newTodo)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(response, "Could not save todo")
		return
	}

}

func ReadTodos(response http.ResponseWriter, request *http.Request) {
	//find the user associted with the jwt
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	todos, err := DbStore.ReadTodos(user.Email)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(response, "Could not read todos")
		return
	}

	json.NewEncoder(response).Encode(todos)
}
