package app

import (
	"context"
	"testing"
	"vrnvgasu/anti-bruteforce/internal/app/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestClear(t *testing.T) {
	const (
		login = "login"
		ip    = "192.168.5.1"
	)

	storageMock := mocks.NewStorage(t)
	storageMock.EXPECT().DeleteByKey(mock.Anything, ip).Return(nil)
	storageMock.EXPECT().DeleteByKey(mock.Anything, login).Return(nil)

	a := App{storage: storageMock}
	err := a.Clear(context.TODO(), ClearRequest{
		Login: login,
		IP:    ip,
	})
	require.NoError(t, err)
}
