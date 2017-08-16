package server

import (
	"encoding/json"
	"lista/api/util"
	"net/http"
)

func profile(response http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("USER-CLAIM").(util.UserClaims)
	json.NewEncoder(response).Encode(struct{ Email string }{user.Email})
}
