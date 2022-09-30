package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	log.Println(generateTokenWithCustomClaim())

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiaGVsbG9vbyB3b3JsZCIsInBhc3N3b3JkIjoiIiwiYXVkIjoibWVlIiwiZXhwIjoxNjY0NTEzOTc4LCJqdGkiOiIxIiwiaXNzIjoiMTIiLCJzdWIiOiJ0ZXN0ICJ9.Vc3TY1XVO_DhPFujBCxVTa0fqK2pe1JGWM_ucb24_BA"

	JwtVerify(token)

	// GeneratehashPassword("1")
}

func generateTokenWithCustomClaim() (string, error) {

	mySignedKey := []byte("AllYourBase")

	type MyCustomClaim struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		jwt.StandardClaims
	}

	claims := MyCustomClaim{
		Name: "hellooo world",
		StandardClaims: jwt.StandardClaims{
			Audience:  "mee",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Id:        "1",
			IssuedAt:  0,
			Issuer:    "12",
			NotBefore: 0,
			Subject:   "test ",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	string, err := token.SignedString(mySignedKey)
	if err != nil {
		return "not happened ", err
	}

	return string, nil
}

func JwtVerify(token string) string {
	tokenn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		log.Println(err)
		return "token failed"
	}

	if tokenn.Valid {
		log.Println("allow")
		return "allow"
	} else {
		log.Println("not allowed")
	}

	log.Println()
	return "unauthorized user"

}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	log.Println(string(bytes))

	return string(bytes), err

}
