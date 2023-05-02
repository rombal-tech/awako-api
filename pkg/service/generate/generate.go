package generate

import (
	"crypto/sha1"
	"fmt"
	"github.com/execaus/exloggo"
	"github.com/sethvargo/go-password/password"
	"math/rand"
	"os"
	"time"
)

func Password() string {
	var err error
	rand.Seed(time.Now().UnixNano())
	inputPassword, err := password.Generate(10, 5, 0, false, true)
	if err != nil {
		exloggo.Fatal(err.Error())
	}
	return inputPassword
}

func Session() string {
	inputString, err := password.Generate(60, 5, 0, false, true)
	if err != nil {
		exloggo.Fatal(err.Error())
	}
	return inputString
}

func PasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
