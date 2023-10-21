package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"github.com/epessine/go-products-server/application/factory"
	"github.com/epessine/go-products-server/application/grpc/pb"
)

func Start(database *gorm.DB, port int) {
	server := grpc.NewServer()
	reflection.Register(server)

	productService := NewProductService(factory.ProductUseCase(database))

	pb.RegisterProductServiceServer(server, productService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error starting gRPC server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("Error starting gRPC server", err)
	}
}
