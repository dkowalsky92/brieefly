package crypto

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/dkowalsky/brieefly/log"
)

// Hash -
func Hash(value string) (*string, error) {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	val := string(hash)
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return &val, nil
}

// CompareHash -
func CompareHash(plain, hash string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hash)
	bytePlain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}
