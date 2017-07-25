package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//source https://github.com/dgrijalva/jwt-go/blob/master/http_example_test.go

type UserClaims struct {
	*jwt.StandardClaims
	Email string
}

//CreateToken will generate a jwt token with the a claim for
//the email
func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &UserClaims{
		&jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 10).Unix()},
		email}
	return token.SignedString([]byte("foobar"))
}
