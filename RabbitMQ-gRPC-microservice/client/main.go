package main

import (
	"context"
	"log"
	"sync"

	pb "gihub.com/leebrouse/RabbitMQ-gRPC-microservice/client/microservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 连接到 gRPC 服务器
func NewPaymentServiceClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50031", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to PaymentService: %v", err)
	}
	return conn, err
}

func NewShoppingServiceClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50032", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to ShoppingService: %v", err)
	}
	return conn, err
}

func NewOrderServiceClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}
	return conn, err
}

// Test Set
func ProcessPaymentTest(wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时，减少 WaitGroup 计数

	paymentConn, err := NewPaymentServiceClient()
	if err != nil {
		return
	}
	defer paymentConn.Close()

	paymentClient := pb.NewPaymentServiceClient(paymentConn)

	resp, err := paymentClient.ProcessPayment(context.Background(), &pb.PaymentRequest{
		UserId: "10",
		Amount: 1000,
	})
	if err != nil {
		log.Printf("PaymentService error: %v", err)
		return
	}
	log.Printf("Payment response: %v", resp)
}

func AddToCartTest(wg *sync.WaitGroup) {
	defer wg.Done()

	shoppingConn, err := NewShoppingServiceClient()
	if err != nil {
		return
	}
	defer shoppingConn.Close()

	shoppingClient := pb.NewShoppingServiceClient(shoppingConn)

	resp, err := shoppingClient.AddToCart(context.Background(), &pb.AddToCartRequest{
		UserId:    "10",
		ProductId: "33",
		Quantity:  30,
	})
	if err != nil {
		log.Printf("ShoppingService error (AddToCart): %v", err)
		return
	}
	log.Printf("AddToCart response: %v", resp)
}

func CheckoutTest(wg *sync.WaitGroup) {
	defer wg.Done()

	shoppingConn, err := NewShoppingServiceClient()
	if err != nil {
		return
	}
	defer shoppingConn.Close()

	shoppingClient := pb.NewShoppingServiceClient(shoppingConn)

	resp, err := shoppingClient.Checkout(context.Background(), &pb.CheckoutRequest{
		UserId: "10",
	})
	if err != nil {
		log.Printf("ShoppingService error (Checkout): %v", err)
		return
	}
	log.Printf("Checkout response: %v", resp)
}

func CreateOrderTest(wg *sync.WaitGroup) {
	defer wg.Done()

	orderConn, err := NewOrderServiceClient()
	if err != nil {
		return
	}
	defer orderConn.Close()

	orderClient := pb.NewOrderServiceClient(orderConn)

	resp, err := orderClient.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		UserId:     "10",
		ProductIds: []string{"33", "11"},
	})
	if err != nil {
		log.Printf("OrderService error: %v", err)
		return
	}
	log.Printf("CreateOrder response: %v", resp)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4) // 添加 4 个 Goroutine 计数

	// 运行所有测试
	go ProcessPaymentTest(&wg)
	go AddToCartTest(&wg)
	go CheckoutTest(&wg)
	go CreateOrderTest(&wg)

	wg.Wait() // 等待所有 Goroutine 完成
	log.Println("All gRPC requests completed.")
}
