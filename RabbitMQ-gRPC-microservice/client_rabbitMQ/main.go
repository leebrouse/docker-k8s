package main

import (
	"context"
	"log"
	"time"

	pb "gihub.com/leebrouse/RabbitMQ-gRPC-microservice/client_rabbitMQ/microservice"
	"google.golang.org/grpc"
)

func connectGRPC(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func AddToCartTest(client pb.ShoppingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.AddToCart(ctx, &pb.AddToCartRequest{
		UserId:    "10",
		ProductId: "33",
		Quantity:  30,
	})
	if err != nil {
		log.Printf("Failed to add to cart: %v", err)
		return
	}
	log.Printf("AddToCart response: %+v", resp)
}

func CheckoutTest(client pb.ShoppingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Checkout(ctx, &pb.CheckoutRequest{UserId: "10"})
	if err != nil {
		log.Printf("Failed to checkout: %v", err)
		return
	}
	log.Printf("Checkout response: %+v", resp)
}

func ProcessPaymentTest(client pb.PaymentServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.ProcessPayment(ctx, &pb.PaymentRequest{
		UserId: "10",
		Amount: 1000,
	})
	if err != nil {
		log.Printf("Failed to process payment: %v", err)
		return
	}
	log.Printf("Payment response: %+v", resp)
}

func CreateOrderTest(client pb.OrderServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
		UserId:     "10",
		ProductIds: []string{"33", "11"},
	})
	if err != nil {
		log.Printf("Failed to create order: %v", err)
		return
	}
	log.Printf("CreateOrder response: %+v", resp)
}

func main() {
	shoppingConn, err := connectGRPC("localhost:50033")
	if err != nil {
		log.Fatalf("Failed to connect to shopping service: %v", err)
	}
	defer shoppingConn.Close()
	shoppingClient := pb.NewShoppingServiceClient(shoppingConn)

	AddToCartTest(shoppingClient)
	CheckoutTest(shoppingClient)
	// ProcessPaymentTest(paymentClient)
	// CreateOrderTest(orderClient)
	log.Println("All gRPC requests completed.")
}