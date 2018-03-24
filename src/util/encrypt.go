package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

/*GetBcryptHash util
GetBcryptHash
@param pw the password you need to encrypt
*/
func GetBcryptHash(pw string) string {
	// generate the bcrypt_pw
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}

/*BcryptAuth util
bcrypt_password authentification
@param captain [description]
*/
func BcryptAuth(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
