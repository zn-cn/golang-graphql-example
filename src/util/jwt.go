// Package util implements some simple tools function for the project
package util

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims util
type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// CreateJWTToken util
func CreateJWTToken(secretKey, method string, claims jwt.Claims) (string, error) {
	// Embed User information to `token`
	token := jwt.NewWithClaims(jwt.GetSigningMethod(method), claims)
	// token -> string. Only server knows this secret (foobar).
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenString, err
}

// ValidateJWTToken util
func ValidateJWTToken(tokenString, secretKey string, claims *CustomClaims) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return token.Valid, err
}
