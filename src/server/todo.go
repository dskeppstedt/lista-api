package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lista/api/models"
	"lista/api/util"
	"log"
	"net/http"
	"strings"
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

func ChangeTodo(response http.ResponseWriter, request *http.Request) {
	if len(strings.Split(request.URL.Path, "/")) < 3 {
		fmt.Fprintln(response, "Malformed url, to long")
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	id := strings.Split(request.URL.Path, "/")[2]
	log.Println(id)

	switch request.Method {
	case "DELETE":
		deleteTodo(response, request, id)
	case "PUT":
		updateTodo(response, request, id)
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request, id string) {
	log.Println("Delete todo with id", id)
	user := r.Context().Value("USER-CLAIM").(util.UserClaims)
	err := DbStore.DeleteTodo(user.Email, id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Could not delet todo with id")
		return
	}

}

func updateTodo(response http.ResponseWriter, request *http.Request, id string) {
	log.Println("Update todo with id", id)

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

	log.Println(requestTodo)
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	err := DbStore.UpdateTodo(user.Email, id, requestTodo)

	if err != nil {
		fmt.Fprintln(response, "Could not find todo to update")
		response.WriteHeader(http.StatusNotFound)
		return
	}
}
