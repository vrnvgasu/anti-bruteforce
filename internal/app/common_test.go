package app

import (
	"testing"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"github.com/stretchr/testify/require"
)

func TestSubnetToString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		network     *pb.Subnet
		expected    string
		expectedErr require.ErrorAssertionFunc
	}{
		{
			name: "success",
			network: &pb.Subnet{
				Ip:   "192.168.5.0",
				Mask: "24",
			},
			expected:    "192.168.5.0/24",
			expectedErr: require.NoError,
		},
		{
			name: "not valid subnet",
			network: &pb.Subnet{
				Ip:   "192.168.5.0",
				Mask: "dummy",
			},
			expected:    "",
			expectedErr: require.Error,
		},
		{
			name:        "null",
			network:     nil,
			expected:    "",
			expectedErr: require.Error,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := App{}
			res, err := a.subnetToString(tt.network)
			tt.expectedErr(t, err)
			require.Equal(t, tt.expected, res)
		})
	}
}
