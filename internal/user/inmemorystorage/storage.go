package inmemorystorage

import (
	"context"
	"fmt"
	"sync"

	"github.com/maxkuzn/tgbot/internal/user"
)

type storage struct {
	m      sync.RWMutex
	users  map[user.ID]user.User
	lastID user.ID
}

func New() *storage {
	return &storage{
		users:  make(map[user.ID]user.User),
		lastID: 1,
	}
}

func (s *storage) Register(_ context.Context, u user.User) (user.ID, error) {
	s.m.Lock()
	defer s.m.Unlock()

	id := s.lastID
	s.lastID++

	if _, ok := s.users[id]; ok {
		return 0, fmt.Errorf("user with id %d already exists", id)
	}

	u.ID = id
	s.users[id] = u

	return id, nil
}

func (s *storage) Get(_ context.Context, userID user.ID) (user.User, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	u, ok := s.users[userID]
	if !ok {
		return user.User{}, fmt.Errorf("user id %d: %w", userID, user.ErrUserNotFound)
	}

	return u, nil
}
