package server

import (
	"fmt"
	"lista/api/db"
	"lista/api/util"
	"log"
	"net/http"
	"os"
	"time"

	"context"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

var DbStore *db.Mongodb
var origin string

//Start is used to start listening for http requests
//Parameters:
// - port, the port that the conneciton will use
func Start(port string) {
	setupRoutes()

	env, set := os.LookupEnv("LISTA_CORS")
	if !set {
		log.Fatal("CORS is not set, fix that please")
	}
	origin = env

	log.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}

func setupRoutes() {
	http.HandleFunc("/info", timer(cors(appInfo)))
	http.HandleFunc("/signup", timer(cors(post(signup))))
	http.HandleFunc("/auth", timer(cors(auth)))
	http.HandleFunc("/profile", timer(cors(protected(profile))))
	http.HandleFunc("/todo", timer(cors(protected(post(CreateTodo)))))
	http.HandleFunc("/todos", timer(cors(protected(ReadTodos))))
	//handles both delete and update, will dispatch accordingly to delete/updateTodo
	http.HandleFunc("/todo/", timer(cors(protected(ChangeTodo))))

}

//MIDDLEWARE

func cors(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		response.Header().Set("Access-Control-Allow-Origin", origin)
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		response.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,UPDATE")

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
