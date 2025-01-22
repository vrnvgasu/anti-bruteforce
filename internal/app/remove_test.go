package app

import (
	"context"
	"testing"
	"vrnvgasu/anti-bruteforce/internal/app/mocks"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRemoveFromWhiteList(t *testing.T) {
	const (
		ip      = "192.168.5.0"
		mask    = "24"
		network = "192.168.5.0/24"
	)

	storageMock := mocks.NewStorage(t)
	storageMock.EXPECT().RemoveFromSet(mock.Anything, pb.ListType_WHITE.String(), network).Return(nil)

	a := App{storage: storageMock}
	err := a.RemoveFromWhiteList(context.TODO(), RemoveFromListRequest{
		ListType: pb.ListType_WHITE,
		Subnet: &pb.Subnet{
			Ip:   ip,
			Mask: mask,
		},
	})
	require.NoError(t, err)
}
