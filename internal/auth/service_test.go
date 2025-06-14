package auth_test

import (
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"testing"
)

type MockUserRepo struct{}

func (repo *MockUserRepo) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}

func (repo *MockUserRepo) GetEmail(email string) (*user.User, error) {
	return nil, nil
}
func TestRegistertest(t *testing.T) {
	const initemail = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepo{})

	email, err := authService.Register(initemail, "1", "Asan")

	if err != nil {
		t.Fatal(err)
	}

	if email != initemail {
		t.Fatal(email)
	}
}
