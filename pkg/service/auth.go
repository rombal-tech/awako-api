package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Registration
}

func NewAuthService(repo repository.Registration) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.Account) (string, error) {
	if user.Password == "" {
		user.Password = generatePassword(user.Password)
		user.Password = generatePasswordHash(user.Password)
		return s.repo.CreateUser(user)
	} else {
		user.Password = generatePasswordHash(user.Password)
		return s.repo.CreateUser(user)
	}
}

func (s *AuthService) CreateSession(sessionDB models.Session, email, password string) (string, error) {
	var err error
	sessionDB.Email, err = s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	sessionDB.SessionString = generateSession(sessionDB.SessionString)
	return s.repo.CreateSession(sessionDB)

}

func (s *AuthService) CreateScheme(schema models.Scheme) (int64, error) {
	return s.repo.CreateSchema(schema)
}

func (s *AuthService) AuthorizationСheck(hed string) bool {
	return s.repo.AuthorizationСheck(hed)
}

func generatePassword(inputPassword string) string {
	var err error
	rand.Seed(time.Now().UnixNano())
	inputPassword, err = password.Generate(10, 5, 0, false, true)
	if err != nil {
		logrus.Fatal(err)
	}
	return inputPassword
}

func generateSession(inputString string) string {
	var err error
	inputString, err = password.Generate(60, 5, 0, false, true)
	if err != nil {
		logrus.Fatal(err)
	}
	return inputString
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
