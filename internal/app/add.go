package app

import (
	"context"
	"fmt"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"
)

type AddToListRequest struct {
	ListType pb.ListType `json:"listType"`
	Subnet   *pb.Subnet  `json:"subnet"`
}

func (a *App) AddToList(ctx context.Context, req AddToListRequest) error {
	network, err := a.subnetToString(req.Subnet)
	if err != nil {
		return fmt.Errorf("app AddToList subnetToString: %w", err)
	}

	if err = a.storage.AddToSet(ctx, req.ListType.String(), network); err != nil {
		return fmt.Errorf("app AddToList AddToSet: %w", err)
	}

	return nil
}
