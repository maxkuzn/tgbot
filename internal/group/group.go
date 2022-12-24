package group

import (
	"time"

	"github.com/maxkuzn/tgbot/internal/user"
)

type ID uint64

type Group struct {
	ID      ID
	Name    string
	AdminID user.ID
	Users   map[user.ID]UserInfo
}

type UserInfo struct {
	UserID   user.ID
	JoinTime time.Time
	// TODO: Add exit time
}
