package inmemorystorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/maxkuzn/tgbot/internal/user"
)

func TestStorage_Get(t *testing.T) {
	testCases := []struct {
		name         string
		createBefore int
		createAfter  int
		adminID      user.ID
		groupName    string
	}{
		{
			name:      "first",
			adminID:   2,
			groupName: "test",
		},
		{
			name:         "several groups",
			createBefore: 20,
			createAfter:  10,
			adminID:      7,
			groupName:    "several",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			s := New()

			for i := 0; i < tc.createBefore; i++ {
				_, err := s.Create(ctx, 1, "-")
				require.NoError(t, err)
			}

			groupID, err := s.Create(ctx, tc.adminID, tc.groupName)
			require.NoError(t, err)

			for i := 0; i < tc.createAfter; i++ {
				_, err := s.Create(ctx, 1, "-")
				require.NoError(t, err)
			}

			g, err := s.Get(ctx, groupID)
			require.NoError(t, err)

			assert.Equal(t, groupID, g.ID)
			assert.Equal(t, tc.adminID, g.AdminID)
			assert.Equal(t, tc.groupName, g.Name)
		})
	}
}
