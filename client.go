package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"userapp/userpb"
)

func main() {
	// 连接 gRPC 服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发送请求
	resp, err := client.RegisterUser(ctx, &userpb.RegisterRequest{
		Name:  "Alice",
		Email: "alice@example.com",
	})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	log.Printf("返回结果: success=%v, message=%s", resp.Success, resp.Message)
}
