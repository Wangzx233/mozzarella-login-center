package main

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"log"
	"mozzarella-login-center/rpc"
)

func main() {
	rpc.RegisterCenter()
	token, err := rpc.KeyCenter.CreateToken(context.Background(), &rpc.CreateTokenReq{
		Domain: "xcx",
		Uid:    uuid.NewV4().String(),
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(token)
}
