package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-login-center/resps"
	"mozzarella-login-center/rpc"
)

func GetTestToken(c *gin.Context) {
	resp, err := rpc.KeyCenter.CreateToken(context.Background(), &rpc.CreateTokenReq{
		Domain: "xcx",
		Uid:    "test_uid",
	})
	if err != nil {
		log.Println(err)
		resps.ParamError(c)
		return
	}

	resps.OKWithData(c, gin.H{
		"access_token":  resp.Token,
		"refresh_token": resp.RefreshToken,
	})
}
