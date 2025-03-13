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

type shoppingServer struct {
	pb.UnimplementedShoppingServiceServer
	conn *amqp.Connection
}

func (s *shoppingServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	// Simulate adding to cart
	fmt.Println(req)
	return &pb.AddToCartResponse{Success: true}, nil
}

func (s *shoppingServer) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	// Simulate checkout
	fmt.Println(req)
	orderID := fmt.Sprintf("order_%s", req.UserId)
	return &pb.CheckoutResponse{Success: true, OrderId: orderID}, nil
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	//newgrpc
	grpcServer := grpc.NewServer()
	pb.RegisterShoppingServiceServer(grpcServer, &shoppingServer{conn: conn})

	listener, err := net.Listen("tcp", ":50032")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	log.Println("Shopping service is running on port 50032")
	
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
