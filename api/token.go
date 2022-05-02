package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-login-center/resps"
	"mozzarella-login-center/rpc"
)

func Refresh(c *gin.Context) {
	refreshToken := c.Query("refresh_token")

	token, err := rpc.KeyCenter.RefreshToken(context.Background(), &rpc.RefreshTokenReq{Rt: refreshToken})
	if err != nil {
		log.Println(err)
		resps.VerifyError(c)
		return
	}

	resps.OKWithData(c, gin.H{
		"access_token":  token.Token,
		"refresh_token": token.RefreshToken,
		"expired_at":    token.ExpiredAt,
	})
}
