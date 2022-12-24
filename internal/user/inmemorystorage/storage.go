package inmemorystorage

import (
	"fmt"

	"github.com/maxkuzn/tgbot/internal/user"
)

type storage struct {
	users  map[user.ID]user.User
	lastID user.ID
}

func New() *storage {
	return &storage{
		users:  make(map[user.ID]user.User),
		lastID: 1,
	}
}

func (s *storage) Register(u user.User) (user.User, error) {
	id := s.lastID
	s.lastID++

	if _, ok := s.users[id]; ok {
		return user.User{}, fmt.Errorf("user with id %d already exists", id)
	}

	u.ID = id
	s.users[id] = u

	return u, nil
}

func (s *storage) Get(userID user.ID) (user.User, error) {
	u, ok := s.users[userID]
	if !ok {
		return user.User{}, fmt.Errorf("user id %d: %w", userID, user.ErrUserNotFound)
	}

	return u, nil
}
