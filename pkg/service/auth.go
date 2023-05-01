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

func (s *AuthService) CreateUser(user models.AccountInput) (*models.AccountRegistrationOutput, error) {
	if user.Password == "" {
		user.Password = GeneratePassword(user.Password)
		user.Password = GeneratePasswordHash(user.Password)
		return s.repo.CreateUser(user)
	} else {
		user.Password = GeneratePasswordHash(user.Password)
		return s.repo.CreateUser(user)
	}
}

func (s *AuthService) CreateSession(sessionDB *models.Session, email, password string) (*models.SessionOutput, error) {
	var err error
	sessionDB.Email, err = s.repo.GetUser(email, GeneratePasswordHash(password))
	if err != nil {
		return nil, err
	}
	sessionDB.SessionString = GenerateSession(sessionDB.SessionString)
	return s.repo.CreateSession(sessionDB)

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

func GeneratePassword(inputPassword string) string {
	var err error
	rand.Seed(time.Now().UnixNano())
	inputPassword, err = password.Generate(10, 5, 0, false, true)
	if err != nil {
		exloggo.Fatal(err.Error())
	}
	return inputPassword
}

func GenerateSession(inputString string) string {
	var err error
	inputString, err = password.Generate(60, 5, 0, false, true)
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
