package server

import (
	"fmt"
	"lista/api/util"
	"net/http"
)

func profile(response http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	fmt.Fprintln(response, "Welcome", user.Email)
}
