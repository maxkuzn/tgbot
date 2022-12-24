package inmemorystorage

import (
	"context"
	"fmt"
	"sync"

	"github.com/maxkuzn/tgbot/internal/group"
	"github.com/maxkuzn/tgbot/internal/user"
)

type storage struct {
	m      sync.RWMutex
	groups map[group.ID]group.Group
	lastID group.ID
}

func New() *storage {
	return &storage{
		groups: make(map[group.ID]group.Group),
		lastID: 1,
	}
}

func (s *storage) Create(_ context.Context, adminID user.ID, name string) (group.ID, error) {
	s.m.Lock()
	defer s.m.Unlock()

	id := s.lastID
	s.lastID++

	if _, ok := s.groups[id]; ok {
		return 0, fmt.Errorf("group %d already exists", id)
	}

	g := group.Group{
		ID:      id,
		Name:    name,
		AdminID: adminID,
		Users:   nil, // TODO: get timestamp?
	}

	s.groups[id] = g

	return id, nil
}

func (s *storage) AddUser(_ context.Context, groupID group.ID, userID user.ID) error {
	s.m.Lock()
	defer s.m.Unlock()

	// TODO: get timestamp?
	return nil
}

func (s *storage) Get(_ context.Context, groupID group.ID) (group.Group, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	g, ok := s.groups[groupID]
	if !ok {
		return group.Group{}, group.ErrGroupNotFound
	}

	return g, nil
}
