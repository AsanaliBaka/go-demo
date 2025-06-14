package di

import (
	"go/adv-demo/internal/user"
)

type IStatRepo interface {
	AddClic(linkId uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	GetEmail(email string) (*user.User, error)
}
