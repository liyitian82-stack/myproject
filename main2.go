package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"userapp/userpb"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 数据库模型
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

// gRPC 服务实现
type userServer struct {
	userpb.UnimplementedUserServiceServer
	db *gorm.DB
}

// RegisterUser 实现
func (s *userServer) RegisterUser(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return &userpb.RegisterResponse{
			Success: false,
			Message: "注册失败",
		}, err
	}

	log.Printf("注册用户成功: name=%s, email=%s", user.Name, user.Email)
	return &userpb.RegisterResponse{
		Success: true,
		Message: fmt.Sprintf("用户 %s 注册成功！", user.Name),
	}, nil
}

func main() {
	// 连接 PostgreSQL
	dsn := "host=localhost user=userapp password=password123 dbname=userdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移用户表
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("迁移数据库失败: %v", err)
	}

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userServer{db: db})

	log.Println("🚀 gRPC 服务已启动，监听端口 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动失败: %v", err)
	}
}
