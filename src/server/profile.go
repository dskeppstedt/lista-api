package server

import (
	"fmt"
	"net/http"
)

func profile(response http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("USER-CLAIM").(UserClaims)
	fmt.Fprintln(response, "Welcome", user.Email)
}
