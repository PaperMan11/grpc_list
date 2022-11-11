package main

import (
	"context"
	"fmt"
	"user/discovery"

	"user/internal/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var etcdAddrs = []string{"127.0.0.1:2379"}

func main() {
	r := discovery.NewResolver(etcdAddrs, zap.NewNop())
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+":///userService", grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to dial %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	resp1, err := c.UserRegister(context.Background(), &pb.RegRequest{UserName: "tan", UserPassword: "123456", UserPassword2: "123456"})
	fmt.Println(resp1, err)
	id := resp1.UserDetail.UserId
	resp, err := c.UserLogin(context.Background(), &pb.LoginRequest{UserId: id, UserPassword: "123456"})
	fmt.Println(resp, err)

}
