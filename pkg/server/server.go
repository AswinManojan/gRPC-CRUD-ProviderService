package server

import (
	"fmt"
	"log"
	"net"

	"github.com/grpc/gobus/service/config"
	handler "github.com/grpc/gobus/service/pkg/Handler"
	pb "github.com/grpc/gobus/service/pkg/PB"
	"google.golang.org/grpc"
)

func NewGRPCServer(cfg *config.Config, handler *handler.ProviderHandler) error {
	log.Println("Connecting to gRPC Server")
	lis, err := net.Listen("tcp", ":"+cfg.GRPCSERVICEPORT)
	grpcServer := grpc.NewServer()
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	pb.RegisterProviderServicesServer(grpcServer, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	fmt.Println("connected")
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	fmt.Println("heeeeee")
	return nil
}
