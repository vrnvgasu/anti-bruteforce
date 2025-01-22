package app

import (
	"context"
	"testing"
	"vrnvgasu/anti-bruteforce/internal/app/mocks"
	"vrnvgasu/anti-bruteforce/internal/core/bucket"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCheck(t *testing.T) {
	const (
		login    = "login"
		password = "pass"
		ip       = "192.168.5.1"
	)

	t.Run("success", func(t *testing.T) { //nolint:dupl
		storageMock := mocks.NewStorage(t)
		storageMock.EXPECT().GetSet(mock.Anything, mock.Anything).Return([]string{}, nil).Twice()

		bucketMock := mocks.NewBucket(t)
		bucketMock.EXPECT().Check(mock.Anything, login, bucket.Login).Return(true, nil)
		bucketMock.EXPECT().Check(mock.Anything, password, bucket.Password).Return(true, nil)
		bucketMock.EXPECT().Check(mock.Anything, ip, bucket.IP).Return(true, nil)

		a := App{
			storage: storageMock,
			bucket:  bucketMock,
		}
		res, err := a.Check(context.TODO(), CheckRequest{
			Login:    login,
			Password: password,
			IP:       ip,
		})
		require.NoError(t, err)
		require.True(t, res)
	})

	t.Run("belongs to black list", func(t *testing.T) {
		storageMock := mocks.NewStorage(t)
		storageMock.EXPECT().GetSet(mock.Anything, pb.ListType_BLACK.String()).
			Return([]string{"192.168.5.0/24"}, nil)

		a := App{
			storage: storageMock,
		}
		res, err := a.Check(context.TODO(), CheckRequest{
			Login:    login,
			Password: password,
			IP:       ip,
		})
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("belongs to white list", func(t *testing.T) {
		storageMock := mocks.NewStorage(t)
		storageMock.EXPECT().GetSet(mock.Anything, pb.ListType_BLACK.String()).Return([]string{}, nil)
		storageMock.EXPECT().GetSet(mock.Anything, pb.ListType_WHITE.String()).
			Return([]string{"192.168.5.0/24"}, nil)

		a := App{
			storage: storageMock,
		}
		res, err := a.Check(context.TODO(), CheckRequest{
			Login:    login,
			Password: password,
			IP:       ip,
		})
		require.NoError(t, err)
		require.True(t, res)
	})

	t.Run("login failed", func(t *testing.T) {
		storageMock := mocks.NewStorage(t)
		storageMock.EXPECT().GetSet(mock.Anything, mock.Anything).Return([]string{}, nil).Twice()

		bucketMock := mocks.NewBucket(t)
		bucketMock.EXPECT().Check(mock.Anything, login, bucket.Login).Return(false, nil)

		a := App{
			storage: storageMock,
			bucket:  bucketMock,
		}
		res, err := a.Check(context.TODO(), CheckRequest{
			Login:    login,
			Password: password,
			IP:       ip,
		})
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("password failed", func(t *testing.T) {
		storageMock := mocks.NewStorage(t)
		storageMock.EXPECT().GetSet(mock.Anything, mock.Anything).Return([]string{}, nil).Twice()

		bucketMock := mocks.NewBucket(t)
		bucketMock.EXPECT().Check(mock.Anything, login, bucket.Login).Return(true, nil)
		bucketMock.EXPECT().Check(mock.Anything, password, bucket.Password).Return(false, nil)

		a := App{
			storage: storageMock,
			bucket:  bucketMock,
		}
		res, err := a.Check(context.TODO(), CheckRequest{
			Login:    login,
			Password: password,
			IP:       ip,
		})
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("ip failed", func(t *testing.T) { //nolint:dupl
		storageMock := mocks.NewStorage(t)
		storageMock.EXPECT().GetSet(mock.Anything, mock.Anything).Return([]string{}, nil).Twice()

		bucketMock := mocks.NewBucket(t)
		bucketMock.EXPECT().Check(mock.Anything, login, bucket.Login).Return(true, nil)
		bucketMock.EXPECT().Check(mock.Anything, password, bucket.Password).Return(true, nil)
		bucketMock.EXPECT().Check(mock.Anything, ip, bucket.IP).Return(false, nil)

		a := App{
			storage: storageMock,
			bucket:  bucketMock,
		}
		res, err := a.Check(context.TODO(), CheckRequest{
			Login:    login,
			Password: password,
			IP:       ip,
		})
		require.NoError(t, err)
		require.False(t, res)
	})
}
