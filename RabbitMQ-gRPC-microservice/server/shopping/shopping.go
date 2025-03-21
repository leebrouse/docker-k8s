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

// const replyQueue = "rpc_response_queue"

type shoppingServer struct {
	pb.UnimplementedShoppingServiceServer
	conn *amqp.Connection
}

func (s *shoppingServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	fmt.Println("Adding to cart:", req)
	//TODO: function for adding to the shopping database
	return &pb.AddToCartResponse{Success: true}, nil
}

func (s *shoppingServer) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	fmt.Println("Checkout request received:", req)

	// 将 CheckoutRequest 发送到 RabbitMQ 的 payment 队列
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"payment_queue", 
		false, 
		false, 
		false, 
		false, 
		nil,
	)
	if err != nil {
		return nil, err
	}

	body, _ := json.Marshal(req)
	err = ch.Publish(
		"",
		q.Name, 
		false, 
		false, 
		amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CheckoutResponse{Success: true, OrderId: fmt.Sprintf("order_%s", req.UserId)}, nil
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterShoppingServiceServer(grpcServer, &shoppingServer{conn: conn})

	lis, err := net.Listen("tcp", ":50033")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Shopping service is running on port 50033")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
