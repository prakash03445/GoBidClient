package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/prakash03445/GoBidProto/gen/go"
	"google.golang.org/grpc"
)

type GoBidClient struct {
	conn   *grpc.ClientConn
	client pb.AuctionServiceClient
}

func NewGoBidClient(serverAddr string) *GoBidClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := pb.NewAuctionServiceClient(conn)
	return &GoBidClient{
		conn:   conn,
		client: client,
	}
}

func (c *GoBidClient) Close() {
	c.conn.Close()
}

func (c *GoBidClient) AddProduct(productID, name, description string, price float64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.Product{
		ProductId:     productID,
		Name:          name,
		Description:   description,
		StartingPrice: price,
	}

	res, err := c.client.AddProduct(ctx, req)
	if err != nil {
		log.Printf("AddProduct error: %v", err)
		return
	}
	fmt.Printf("AddProduct response: %v - %v\n", res.Success, res.Message)
}

func (c *GoBidClient) GetProducts() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := c.client.GetProducts(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("GetProducts error: %v", err)
		return
	}

	fmt.Println("Current Products:")
	for _, p := range res.Product {
		fmt.Printf("- [%s] %s: %s, starting at %.2f\n", p.ProductId, p.Name, p.Description, p.StartingPrice)
	}
}
