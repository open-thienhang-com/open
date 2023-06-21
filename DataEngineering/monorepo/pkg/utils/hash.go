package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
)

var secretKey = "neictr98y85klfgneghre"

// Create a salt string with 32 bytes of crypto/rand data
func generateSalt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// Hash a password with the salt
func hashPassword(plainText string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText+salt)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

// password := checkArgs()
//    salt := generateSalt()
//    hashedPassword := hashPassword(password, salt)
//    fmt.Println("Password: " + password)
//    fmt.Println("Salt: " + salt)
//    fmt.Println("Hashed password: " + hashedPassword)
