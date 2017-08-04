package server

import (
	"fmt"
	"lista/api/db"
	"lista/api/util"
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
	log.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}

func setupRoutes() {
	http.HandleFunc("/info", timer(cors(appInfo)))
	http.HandleFunc("/signup", timer(cors(post(signup))))
	http.HandleFunc("/auth", timer(cors(auth)))
	http.HandleFunc("/profile", timer(cors(protected(profile))))

}

//MIDDLEWARE

func cors(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		//		if origin := request.Header.Get("Origin"); origin != "" {
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		//}

		if request.Method == "OPTIONS" {
			response.WriteHeader(200)
			return
		}

		next.ServeHTTP(response, request)

	})
}

func post(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Only POST is allowed")
			return
		}

		next.ServeHTTP(w, r)

	})
}

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

		var claims util.UserClaims

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
