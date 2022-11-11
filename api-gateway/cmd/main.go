package main

import (
	"api-gateway/config"
	"api-gateway/discovery"
	"api-gateway/pb"
	"api-gateway/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	// conf
	if err := config.Init(); err != nil {
		panic(err)
	}

	// RPC connect
	etcdAddr := []string{config.Conf.EtcdConf.Address}
	etcdResolver := discovery.NewResolver(etcdAddr, zap.NewNop())
	resolver.Register(etcdResolver)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	userSrvClient, err := RPCConnect(ctx, etcdResolver, "userService")
	if err != nil {
		log.Println("rpc connect failed")
		os.Exit(1)
	}
	userSrv := pb.NewUserServiceClient(userSrvClient)

	taskSrvClient, err := RPCConnect(ctx, etcdResolver, "taskService")
	if err != nil {
		log.Println("rpc connect failed")
		os.Exit(1)
	}
	taskSrv := pb.NewTaskServiceClient(taskSrvClient)

	// http server
	r := router.NewRouter(userSrv, taskSrv)
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Conf.ServerConf.Port),
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done
		log.Println("close http server now...")
		server.Shutdown(context.TODO())
	}()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		return
	}
}

func RPCConnect(ctx context.Context, etcdResolver *discovery.Resolver, serviceName string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	addr := fmt.Sprintf("%s:///%s", etcdResolver.Scheme(), serviceName)
	conn, err := grpc.DialContext(ctx, addr, opts...)
	return conn, err
}
