package internalgrpc

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vrnvgasu/anti-bruteforce/internal/config"
	"vrnvgasu/anti-bruteforce/internal/server/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Cmd string

const (
	help                   Cmd = "help"
	checkCmd               Cmd = "check"
	clearCmd               Cmd = "clear"
	addToBlackListCmd      Cmd = "add_to_back_list"
	addToWhiteListCmd      Cmd = "add_to_white_list"
	removeFromBlackListCmd Cmd = "remove_from_back_list"
	removeFromWhiteListCmd Cmd = "remove_from_white_list"
)

const (
	okResult = "OK"
)

type Client struct {
	pb.AntiBruteforceClient
}

func NewClient() (*Client, error) {
	addr := fmt.Sprintf("%s:%d", config.Cfg.GRPSServer.Host, config.Cfg.GRPSServer.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("new grpc client: %w", err)
	}

	client := pb.NewAntiBruteforceClient(conn)

	return &Client{AntiBruteforceClient: client}, nil
}

func (c *Client) Scan() <-chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin) // читаем сообщение из консоли
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if scanner.Err() != nil {
			close(out)
		}
	}()
	return out
}

func (c *Client) ListenCmd(ctx context.Context, stdin <-chan string) {
	for {
		select {
		case request := <-stdin:
			cmd, args := commandAndArgsFromRequest(request)
			res, err := c.exec(ctx, cmd, args...)
			if err != nil {
				res = fmt.Sprintf("get error: %v", err)
			}
			fmt.Printf("get result: %s\n", res)
		case <-ctx.Done():
			return
		}
	}
}

func (c *Client) exec(ctx context.Context, cmd Cmd, args ...string) (string, error) {
	switch cmd {
	case help:
		return `
	help                        List available commands
	----------------------------------------------------------------
	Cmd:                        Example:
	----------------------------------------------------------------
	check                       check login password 192.168.5.1
	clear                       clear login 192.168.5.1
	add_to_back_list            add_to_back_list 192.168.5.0 24
	add_to_white_list           add_to_white_list 192.168.5.0 24
	remove_from_back_list       remove_from_back_list 192.168.5.0 24
	remove_from_white_list      remove_from_white_list 192.168.5.0 24
	`, nil
	case checkCmd:
		if len(args) != 3 {
			return "", fmt.Errorf("check command requires 3 arguments: login, password, ip")
		}

		response, err := c.AntiBruteforceClient.Check(ctx, &pb.CheckRequest{
			Login:    args[0],
			Password: args[1],
			Ip:       args[2],
		})
		if err != nil {
			return "", fmt.Errorf("check command: %w", err)
		}

		return strconv.FormatBool(response.GetResult()), nil
	case clearCmd:
		if len(args) != 2 {
			return "", fmt.Errorf("clear command requires 2 arguments: login, ip")
		}
		_, err := c.AntiBruteforceClient.Clear(ctx, &pb.ClearRequest{
			Login: args[0],
			Ip:    args[1],
		})
		if err != nil {
			return "", fmt.Errorf("clear command: %w", err)
		}

		return okResult, nil
	case addToBlackListCmd:
		if len(args) != 2 {
			return "", fmt.Errorf("add to black list requires 2 arguments: ip, mask")
		}
		_, err := c.AntiBruteforceClient.AddToList(ctx, &pb.ListRequest{
			Subnet: &pb.Subnet{
				Ip:   args[0],
				Mask: args[1],
			},
			Type: pb.ListType_BLACK,
		})
		if err != nil {
			return "", fmt.Errorf("add to black list command: %w", err)
		}

		return okResult, nil
	case addToWhiteListCmd:
		if len(args) != 2 {
			return "", fmt.Errorf("add to white list requires 2 arguments: ip, mask")
		}
		_, err := c.AntiBruteforceClient.AddToList(ctx, &pb.ListRequest{
			Subnet: &pb.Subnet{
				Ip:   args[0],
				Mask: args[1],
			},
			Type: pb.ListType_WHITE,
		})
		if err != nil {
			return "", fmt.Errorf("add to white list command: %w", err)
		}

		return okResult, nil
	case removeFromBlackListCmd:
		if len(args) != 2 {
			return "", fmt.Errorf("remove from black list requires 2 arguments: ip, mask")
		}
		_, err := c.AntiBruteforceClient.RemoveFromList(ctx, &pb.ListRequest{
			Subnet: &pb.Subnet{
				Ip:   args[0],
				Mask: args[1],
			},
			Type: pb.ListType_BLACK,
		})
		if err != nil {
			return "", fmt.Errorf("remove from blacklist command: %w", err)
		}

		return okResult, nil
	case removeFromWhiteListCmd:
		if len(args) != 2 {
			return "", fmt.Errorf("remove from white list requires 2 arguments: ip, mask")
		}
		_, err := c.AntiBruteforceClient.RemoveFromList(ctx, &pb.ListRequest{
			Subnet: &pb.Subnet{
				Ip:   args[0],
				Mask: args[1],
			},
			Type: pb.ListType_WHITE,
		})
		if err != nil {
			return "", fmt.Errorf("remove from white command: %w", err)
		}

		return okResult, nil
	default:
		return "", fmt.Errorf("unknown command: %s", cmd)
	}
}

func commandAndArgsFromRequest(request string) (Cmd, []string) {
	var (
		cmd  Cmd
		args = make([]string, 0)
	)

	items := strings.Split(request, " ")
	if len(items) == 0 {
		return cmd, args
	}

	cmd = Cmd(strings.TrimSpace(items[0]))
	if len(items) == 1 {
		return cmd, args
	}

	for _, arg := range items[1:] {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}

		args = append(args, arg)
	}

	return cmd, args
}
