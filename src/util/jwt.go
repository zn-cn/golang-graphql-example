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
func ValidateJWTToken(tokenString, secretKey string, claims *jwt.Claims) (bool, *CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if err != nil {
		log.Fatal(err)
		return false, nil, err
	}
	return token.Valid, token.Claims.(*CustomClaims), err
}
