package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-login-center/dao"
	"mozzarella-login-center/logic"
	"mozzarella-login-center/model"
	"mozzarella-login-center/rpc"
)

func Register(c *gin.Context) {
	var user model.UserJson
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("bind json err :", err)
		c.JSON(503, gin.H{
			"status": 50000,
			"info":   "param err ",
			"data":   nil,
		})
		return
	}

	//验证电话格式
	err = logic.Check(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 40000,
			"info":   fmt.Sprint(err),
			"data":   nil,
		})
		return
	}

	p, s := logic.IsRegister(user.PhoneNumber, user.StudentID)
	if s || p {
		c.JSON(401, gin.H{
			"status": 50000,
			"info":   "用户已存在",
			"data":   nil,
		})
		return
	}

	//验证code
	code, err := dao.GetCode(user.PhoneNumber)
	if err != nil {
		c.JSON(401, gin.H{
			"status": 40000,
			"info":   "验证码错误",
			"data":   nil,
		})
		return
	}

	if code != user.Code {
		c.JSON(401, gin.H{
			"status": 40000,
			"info":   "验证码错误",
			"data":   nil,
		})
		return
	}

	daoUser := logic.UserJsonToDao(user)
	err = dao.Register(daoUser)
	if err != nil {
		c.JSON(503, gin.H{
			"status": 50000,
			"info":   "err",
			"data":   nil,
		})
		return
	}

	//生成access_token和refresh_token
	resp, err := rpc.KeyCenter.CreateToken(context.Background(), &rpc.CreateTokenReq{
		Domain: "common",
		Uid:    daoUser.Uid,
	})
	if err != nil {
		log.Println(err)
	}

	dao.DelCode(user.PhoneNumber)

	c.JSON(200, gin.H{
		"status": 10000,
		"info":   "success",
		"data": gin.H{
			"access_token":  resp.Token,
			"refresh_token": resp.RefreshToken,
		},
	})
}
