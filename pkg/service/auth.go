package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"log"
	"math/rand"
	"time"
)

const salt = "5hh24h23fgfgg4g3g"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.Account) (string, error) {
	if user.Password == "" {
		user.Password = generatePassword(user.Password)
		//user.Password = generatePasswordHash(user.Password)
		return s.repo.CreateUser(user)
	} else {
		user.Password = generatePasswordHash(user.Password)
		return s.repo.CreateUser(user)
	}
}

func generatePassword(inputPassword string) string {
	var err error
	rand.Seed(time.Now().UnixNano())
	inputPassword, err = password.Generate(10, 5, 0, false, true)
	if err != nil {
		log.Fatal(err)
	}
	return inputPassword
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
