package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func IsMatchHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func main() {
	password := "hahaha"
	hash, _ := HashPassword(password)
	fmt.Println(password)
	fmt.Println(hash)

	fmt.Println(IsMatchHash(password, hash))

}
