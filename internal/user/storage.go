package user

import "errors"

var ErrUserNotFound = errors.New("user not found")

type Storage interface {
	Register(user User) (User, error)
	Get(userID ID) (User, error)
}
