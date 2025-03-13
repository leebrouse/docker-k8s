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

type paymentServer struct {
	pb.UnimplementedPaymentServiceServer
	conn *amqp.Connection
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// Simulate payment processing
	fmt.Println(req)
	transactionID := fmt.Sprintf("txn_%s", req.UserId)
	return &pb.PaymentResponse{Success: true, TransactionId: transactionID}, nil
}

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, &paymentServer{conn: conn})

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50031")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	log.Println("Payment service is running on port 50031")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
