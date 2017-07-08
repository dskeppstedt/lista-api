package server

import (
	"fmt"
	"lista/api/db"
	"log"
	"net/http"
	"time"

	"context"

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
	http.HandleFunc("/info", timer(appInfo))
	http.HandleFunc("/auth", timer(auth))
	http.HandleFunc("/profile", timer(protected(profile)))
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

		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
		})

		log.Println(claims)
		log.Println(token)
		// If the token is missing or invalid, return error
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "You are unauthorized")
			log.Println(w, "Invalid token:", err)

			return
		}

		//add the user claim to the context
		ctx := context.WithValue(r.Context(), "USER-CLAIM", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
