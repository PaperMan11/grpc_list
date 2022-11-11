package main

import (
	"log"
	"net"
	"task/config"
	"task/discovery"
	"task/internal/handler"
	"task/internal/pb"
	"task/internal/repository"
	"task/pkg/snowflake"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	config.Init()
	repository.Init()
	if err := snowflake.Init("2022-10-31", 1); err != nil {
		log.Println("snowflake init failed")
		return
	}
	grpcAddress := config.Conf.ServerConf.GrpcAddress // server address
	// 1. grpc server
	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterTaskServiceServer(server, handler.TaskServer{})
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	// register to etcd
	// etcd address
	var etcdAddress = []string{config.Conf.EtcdConf.Address}
	srvInfo := discovery.Server{
		Name:    "taskService",
		Addr:    grpcAddress,
		Version: "v1",
		Weight:  1,
	}
	r := discovery.NewRegister(etcdAddress, zap.NewNop())
	if _, err := r.Register(srvInfo, 10); err != nil {
		panic(err)
	}

	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
