package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/execaus/exloggo"
	"github.com/sethvargo/go-password/password"
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

func (s *AuthService) CreateAccount(email string, inputPassword string) (*models.Account, error) {
	var hash string
	var accountPassword string

	if inputPassword == "" {
		accountPassword = GeneratePassword()
	} else {
		accountPassword = inputPassword
	}

	hash = GeneratePasswordHash(accountPassword)

	return s.repo.CreateAccount(email, hash)
}

func (s *AuthService) CreateSession(email string) (*models.Session, error) {
	hash := GenerateSession()

	return s.repo.CreateSession(email, hash)
}

func (s *AuthService) CreateScheme(schema models.Scheme, email string) (*models.SchemeOutput, error) {
	return s.repo.CreateScheme(schema, email)
}

func (s *AuthService) CheckAuthorization(hed string) (string, error) {
	return s.repo.CheckAuthorization(hed)
}

func (s *AuthService) GetScheme(email string) ([]models.SchemeOutput, error) {
	return s.repo.GetScheme(email)
}

func GeneratePassword() string {
	var err error
	rand.Seed(time.Now().UnixNano())
	inputPassword, err := password.Generate(10, 5, 0, false, true)
	if err != nil {
		exloggo.Fatal(err.Error())
	}
	return inputPassword
}

func GenerateSession() string {
	inputString, err := password.Generate(60, 5, 0, false, true)
	if err != nil {
		exloggo.Fatal(err.Error())
	}
	return inputString
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
