package bucket

import (
	"context"
	"testing"
	"time"
	storageMock "vrnvgasu/anti-bruteforce/internal/core/bucket/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCheckByStorage(t *testing.T) {
	var (
		ctx                 = context.TODO()
		item                = "test"
		rate  int64         = 10
		limit time.Duration = 60
	)

	t.Run("first increment", func(t *testing.T) {
		st := storageMock.NewStorage(t)
		st.EXPECT().IncrCounterByKey(ctx, item).
			Return(1, nil)
		st.EXPECT().ExpireKey(ctx, item, limit).Return(nil)
		s := Service{
			timeLimit: limit,
			st:        st,
		}
		result, err := s.checkByStorage(ctx, item, rate)
		require.NoError(t, err)
		require.True(t, result)
	})
	t.Run("second increment", func(t *testing.T) {
		st := storageMock.NewStorage(t)
		st.EXPECT().IncrCounterByKey(ctx, item).
			Return(2, nil)
		s := Service{
			timeLimit: limit,
			st:        st,
		}
		result, err := s.checkByStorage(ctx, item, rate)
		require.NoError(t, err)
		require.True(t, result)
	})
	t.Run("to many increment", func(t *testing.T) {
		st := storageMock.NewStorage(t)
		st.EXPECT().IncrCounterByKey(ctx, item).
			Return(rate, nil)
		s := Service{
			timeLimit: limit,
			st:        st,
		}
		result, err := s.checkByStorage(ctx, item, rate)
		require.NoError(t, err)
		require.False(t, result)
	})
}

func TestCheck(t *testing.T) {
	cases := []struct {
		name   string
		entity EntityRate
	}{
		{
			name:   "login",
			entity: Login,
		},
		{
			name:   "password",
			entity: Password,
		},
		{
			name:   "ip",
			entity: IP,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			st := storageMock.NewStorage(t)
			st.EXPECT().IncrCounterByKey(mock.Anything, mock.Anything).
				Return(2, nil)
			s := Service{
				loginRate:    10,
				passwordRate: 10,
				ipRate:       10,
				st:           st,
			}
			result, err := s.Check(context.TODO(), "item", tt.entity)
			require.NoError(t, err)
			require.True(t, result)
		})
	}

	t.Run("error", func(t *testing.T) {
		s := Service{}
		result, err := s.Check(context.TODO(), "dummy", -1)
		require.Error(t, err)
		require.False(t, result)
	})
}
