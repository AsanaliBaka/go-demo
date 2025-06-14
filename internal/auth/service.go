package auth

import (
	"errors"
	"fmt"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo di.IUserRepository
}

func NewAuthService(user di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepo: user,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existed, _ := service.UserRepo.GetEmail(email)

	if existed != nil {
		return "", errors.New(ErrUserExsist)
	}

	hachedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Password: string(hachedPassword),
		Name:     name,
	}

	_, err = service.UserRepo.Create(user)

	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service *AuthService) LoginService(email, password string) (string, error) {
	existed, _ := service.UserRepo.GetEmail(email)

	if existed == nil {
		return "", fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existed.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("password is in correct")
	}

	return existed.Email, nil

}
