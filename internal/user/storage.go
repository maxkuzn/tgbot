package user

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type Storage interface {
	Register(ctx context.Context, user User) (ID, error)
	Get(ctx context.Context, userID ID) (User, error)
}
