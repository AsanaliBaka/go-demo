package user

import "go/adv-demo/pkg/db"

type UserRepo struct {
	db *db.Db
}

func NewUserRepo(db *db.Db) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(user *User) (*User, error) {
	result := u.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserRepo) GetEmail(email string) (*User, error) {
	var newUser User
	result := u.db.First(&newUser, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
