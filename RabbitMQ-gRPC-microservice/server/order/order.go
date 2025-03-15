package main

import (
	"context"
	"encoding/json"
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
	fmt.Println("Creating order:", req)
	orderID := fmt.Sprintf("order_%s", req.UserId)
	return &pb.CreateOrderResponse{Success: true, OrderId: orderID}, nil
}

func (s *orderServer) startListening() {
	ch, err := s.conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("order_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	//for block the server
	var forever chan struct{}
	for msg := range msgs {
		var req pb.PaymentResponse
		json.Unmarshal(msg.Body, &req)
		s.CreateOrder(context.Background(), &pb.CreateOrderRequest{UserId: req.TransactionId})
	}
	<-forever
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	grpcServer := grpc.NewServer()
	server := &orderServer{conn: conn}
	pb.RegisterOrderServiceServer(grpcServer, server)

	go server.startListening()

	lis, err := net.Listen("tcp", ":50032")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Shopping service is running on port 50032")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
