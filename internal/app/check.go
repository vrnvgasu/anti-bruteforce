package app

import (
	"context"
	"fmt"
	"net"
	"vrnvgasu/anti-bruteforce/internal/core/bucket"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"
)

type CheckRequest struct {
	Login    string
	Password string
	IP       string
}

func (a *App) Check(ctx context.Context, req CheckRequest) (ok bool, err error) {
	// черный список
	isBelong, err := a.isIPBelongsToList(ctx, req.IP, pb.ListType_BLACK)
	if err != nil {
		return false, fmt.Errorf("app Check isIPBelongsToList BlackList: %w", err)
	}
	if isBelong {
		return false, nil
	}

	// белый список
	isBelong, err = a.isIPBelongsToList(ctx, req.IP, pb.ListType_WHITE)
	if err != nil {
		return false, fmt.Errorf("app Check isIPBelongsToList WhiteList: %w", err)
	}
	if isBelong {
		return true, nil
	}

	// проверка логина
	ok, err = a.bucket.Check(ctx, req.Login, bucket.Login)
	if err != nil {
		return false, fmt.Errorf("app Check Login bucket.Check: %w", err)
	}
	if !ok {
		return false, nil
	}

	// проверка пароля
	ok, err = a.bucket.Check(ctx, req.Password, bucket.Password)
	if err != nil {
		return false, fmt.Errorf("app Check Password bucket.Check: %w", err)
	}
	if !ok {
		return false, nil
	}

	// проверка ip
	ok, err = a.bucket.Check(ctx, req.IP, bucket.IP)
	if err != nil {
		return false, fmt.Errorf("app Check IP bucket.Check: %w", err)
	}
	if !ok {
		return false, nil
	}

	return true, nil
}

func (a *App) isIPBelongsToList(ctx context.Context, ip string, list pb.ListType) (ok bool, err error) {
	networks, err := a.storage.GetSet(ctx, list.String())
	if err != nil {
		return false, fmt.Errorf("app isIPBelongsToList isIPBelongsToList: %w", err)
	}

	for _, mask := range networks {
		_, subnet, err := net.ParseCIDR(mask)
		if err != nil {
			return false, fmt.Errorf("app isIPBelongsToList ParseCIDR: %w", err)
		}

		ip := net.ParseIP(ip)
		if subnet.Contains(ip) {
			return true, nil
		}
	}

	return false, nil
}
