package util

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//source https://github.com/dgrijalva/jwt-go/blob/master/http_example_test.go

type UserClaims struct {
	*jwt.StandardClaims
	Email string
}

var JwtSecure string

func Init() {
	readJwtSecret()
}

func readJwtSecret() {

	env, set := os.LookupEnv("LISTA_JWT")
	if !set {
		log.Fatal("JWT not set, fix it!")
	}
	JwtSecure = env
}

//CreateToken will generate a jwt token with the a claim for
//the email
func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &UserClaims{
		&jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 10).Unix()},
		email}
	return token.SignedString([]byte(JwtSecure))
}
