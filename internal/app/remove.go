package app

import (
	"context"
	"fmt"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"
)

type RemoveFromListRequest struct {
	ListType pb.ListType `json:"listType"`
	Subnet   *pb.Subnet  `json:"subnet"`
}

func (a *App) RemoveFromWhiteList(ctx context.Context, req RemoveFromListRequest) error {
	network, err := a.subnetToString(req.Subnet)
	if err != nil {
		return fmt.Errorf("app RemoveFromWhiteList subnetToString: %w", err)
	}

	if err = a.storage.RemoveFromSet(ctx, req.ListType.String(), network); err != nil {
		return fmt.Errorf("app RemoveFromWhiteList RemoveFromSet: %w", err)
	}

	return nil
}
