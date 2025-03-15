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

type paymentServer struct {
	pb.UnimplementedPaymentServiceServer
	conn *amqp.Connection
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	fmt.Println("Processing payment:", req)

	transactionID := fmt.Sprintf("txn_%s", req.UserId)

	// 将支付结果放入 RabbitMQ 的 order 队列
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("order_queue", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	response := pb.PaymentResponse{Success: true, TransactionId: transactionID}
	body, _ := json.Marshal(response)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *paymentServer) startListening() {
	ch, err := s.conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("payment_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	var forever chan struct{}
	for msg := range msgs {
		var req pb.PaymentRequest
		json.Unmarshal(msg.Body, &req)
		s.ProcessPayment(context.Background(), &req)
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
	server := &paymentServer{conn: conn}
	pb.RegisterPaymentServiceServer(grpcServer, server)

	go server.startListening()

	lis, err := net.Listen("tcp", ":50031")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Payment service is running on port 50031")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
