package service

import (
	"crypto/sha1" //генерация пароля
	"fmt"
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/SamsonAirapetyan/todo-app/pkg/repository"
)

const salt = "fklasdfl32423xgjdkfg"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

//func (s *AuthService) CreateUser(user todo.User) error {
//	user.Password = generatePasswordHash(user.Password)
//	return s.repo.CreateUser(user)
//}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
