package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gihub.com/leebrouse/RabbitMQ-gRPC-microservice/proto/microservice"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type orderServer struct {
	pb.UnimplementedOrderServiceServer
	conn *amqp.Connection
}

func (s *orderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Simulate order creation
	fmt.Println(req)
	orderID := fmt.Sprintf("order_%s", req.UserId)
	return &pb.CreateOrderResponse{Success: true, OrderId: orderID}, nil
}

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &orderServer{conn: conn})

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50030")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Order service is running on port 50030")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
