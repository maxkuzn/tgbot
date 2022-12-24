package group

import (
	"context"
	"errors"

	"github.com/maxkuzn/tgbot/internal/user"
)

var ErrGroupNotFound = errors.New("group not found")

type Storage interface {
	Create(ctx context.Context, adminID user.ID, name string) (ID, error)
	AddUser(ctx context.Context, groupID ID, userID user.ID) error
	Get(ctx context.Context, groupID ID) (Group, error)
}
