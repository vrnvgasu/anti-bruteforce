package app

import (
	"fmt"
	"net"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"
)

func (a *App) subnetToString(network *pb.Subnet) (string, error) {
	if network == nil {
		return "", fmt.Errorf("app subnetToString is null")
	}

	res := fmt.Sprintf("%s/%s", network.GetIp(), network.GetMask())
	if _, _, err := net.ParseCIDR(res); err != nil {
		return "", fmt.Errorf("app ValidateNetwork ParseCIDR: %w", err)
	}

	return res, nil
}
