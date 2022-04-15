package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

var (
	KeyCenter     MozzarellaBookClient
	KeyCenterConn *grpc.ClientConn
)

func NewKeyCenter() {
	if KeyCenterConn != nil {
		err := KeyCenterConn.Close()
		if err != nil {
			log.Println(err)
		}
	}

	conn, err := grpc.Dial(KeyServiceHost, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
		return
	}

	KeyCenterConn = conn

	KeyCenter = NewMozzarellaBookClient(conn)

}
