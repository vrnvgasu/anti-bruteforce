package app

import (
	"context"
	"testing"
	"vrnvgasu/anti-bruteforce/internal/app/mocks"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddToList(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		list pb.ListType
	}{
		{
			name: "white list",
			list: pb.ListType_WHITE,
		},
		{
			name: "black list",
			list: pb.ListType_BLACK,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storageMock := mocks.NewStorage(t)
			storageMock.EXPECT().AddToSet(mock.Anything, tt.list.String(), mock.Anything).Return(nil)

			a := App{
				storage: storageMock,
			}
			err := a.AddToList(context.TODO(), AddToListRequest{
				ListType: tt.list,
				Subnet: &pb.Subnet{
					Ip:   "192.168.5.0",
					Mask: "24",
				},
			})
			require.NoError(t, err)
		})
	}
}
