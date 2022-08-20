package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/korasdor/todo-app/models"
	"github.com/korasdor/todo-app/repository"
)

const salt = "ksajldfkdjsklf"

type AuthService struct {
	repo repository.Autharization
}

func NewAuthService(repo repository.Autharization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (int, error) {
	//TODO
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
