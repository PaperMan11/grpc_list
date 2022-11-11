package main

import (
	"log"
	"net"
	"user/config"
	"user/discovery"
	"user/internal/controller"
	"user/internal/global"
	"user/internal/pb"
	"user/pkg/snowflake"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	config.Init()
	global.Init()
	if err := snowflake.Init("2022-10-31", 1); err != nil {
		log.Println("snowflake init failed")
		return
	}
	grpcAddress := config.Conf.ServerConf.GrpcAddress

	// 1. grpc server
	server := grpc.NewServer()
	defer server.Stop()
	// bind service
	pb.RegisterUserServiceServer(server, controller.UserService{})
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	// register to etcd
	// etcd address
	var etcdAddress = []string{config.Conf.EtcdConf.Address}

	srvinfo := discovery.Server{
		Name:    "userService",
		Addr:    grpcAddress,
		Version: "v1",
		Weight:  1,
	}
	r := discovery.NewRegister(etcdAddress, zap.NewNop())
	if _, err := r.Register(srvinfo, 10); err != nil {
		panic(err)
	}

	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
