package util

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

//GenerateHash returns a bcrypted string
func GenerateHash(clearText string) string {
	password := []byte(clearText)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err == nil {
		log.Println("The hash and password is a match")
	}
	return string(hashedPassword)
}

func toInt(s string) (int, error) {
	return strconv.Atoi(s)
}
