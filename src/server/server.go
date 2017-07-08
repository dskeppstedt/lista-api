package server

import (
	"fmt"
	"lista/api/db"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

var DbStore *db.Mongodb

//Start is used to start listening for http requests
//Parameters:
// - port, the port that the conneciton will use
func Start(port string) {
	setupRoutes()
	log.Println("Accepting requestons on port", port)
	http.ListenAndServe(port, nil)
}

func setupRoutes() {
	http.HandleFunc("/", timer(hello))
	http.HandleFunc("/info", timer(appInfo))
	http.HandleFunc("/auth", timer(auth))
	http.HandleFunc("/profile", timer(protected(profile)))
}

//HANDLERS
func hello(response http.ResponseWriter, request *http.Request) {

	response.WriteHeader(200)
	response.Write([]byte("Hello!\n"))
}

//MIDDLEWARE
func timer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		before := time.Now()
		next.ServeHTTP(w, r)
		elapsedTime := time.Now().Sub(before)
		log.Println("Request for", r.RequestURI, "delivered in", elapsedTime)
	})
}

func protected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from request

		var claims UserClaims

		//NOTE: A token can be accessed here and it prob
		//should be passed along via context
		_, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
		})


		log.Println(claims)
		// If the token is missing or invalid, return error
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "You are unauthorized")
			log.Println(w, "Invalid token:", err)

			return
		}
		next.ServeHTTP(w, r)
	})
}
