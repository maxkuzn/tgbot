package inmemorystorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/maxkuzn/tgbot/internal/user"
)

func TestStorage_Add(t *testing.T) {
	testCases := []struct {
		name      string
		addBefore []user.User
		get       user.ID
		want      user.User
		wantErr   error
	}{
		{
			name: "one user",
			addBefore: []user.User{
				{Name: "Max"},
			},
			get: 1,
			want: user.User{
				ID:   1,
				Name: "Max",
			},
			wantErr: nil,
		},
		{
			name: "several users",
			addBefore: []user.User{
				{Name: "1"},
				{Name: "2"},
				{Name: "3"},
				{Name: "4"},
			},
			get: 3,
			want: user.User{
				ID:   3,
				Name: "3",
			},
			wantErr: nil,
		},
		{
			name:      "empty",
			addBefore: nil,
			get:       2,
			wantErr:   user.ErrUserNotFound,
		},
		{
			name: "not found",
			addBefore: []user.User{
				{Name: "1"},
				{Name: "2"},
				{Name: "3"},
			},
			get:     4,
			wantErr: user.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			s := New()

			for i, add := range tc.addBefore {
				id, err := s.Register(ctx, add)
				require.NoError(t, err)
				assert.Equal(t, user.ID(i+1), id)
			}

			got, err := s.Get(ctx, tc.get)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
