package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"mozzarella-login-center/dao"
	"net"
)

type LoginCenter struct {
}

func (l *LoginCenter) FindUser(ctx context.Context, request *FindUserRequest) (user *FindUserReply, err error) {
	u, err := dao.FindUserByUid(request.Uid)
	return &FindUserReply{
		StudentID: u.StudentID,
		RealName:  u.RealName,
	}, err
}

func InitRpc() {
	// 监听本地的8901端口
	lis, err := net.Listen("tcp", ":8902")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器

	RegisterMozzarellaLoginCenterServer(s, &LoginCenter{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
