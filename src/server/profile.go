package server

import (
	"fmt"
	"net/http"
)

func profile(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Welcome to your profile!")
}
