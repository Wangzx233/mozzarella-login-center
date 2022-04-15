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

func SendLoginCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	//判断是否存在用户电话号
	p, _ := logic.IsRegister(phoneNumber, "")
	if !p {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "phoneNumber don't exited",
		})
		return
	}

	code, err := logic.SendCode(phoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "phoneNumber don't exited",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   code,
	})
}

func SendRegisterCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	//判断是否存在用户电话号
	p, _ := logic.IsRegister(phoneNumber, "")
	if p {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "phoneNumber exited",
		})
		return
	}

	code, err := logic.SendCode(phoneNumber)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "phoneNumber don't exited",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   code,
	})
}

// LoginByCode 验证码登陆
func LoginByCode(c *gin.Context) {
	var user model.UserJson
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 40000,
			"info":   "parameter mismatch",
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

	//判断是否存在用户电话号
	p, _ := logic.IsRegister(user.PhoneNumber, user.StudentID)
	if !p {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "phoneNumber don't exited",
			"data":   nil,
		})
		return
	}

	err = logic.LoginByCode(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   fmt.Sprint(err),
			"data":   nil,
		})
		return
	}

	u, err := dao.FindUserByPhoneNumber(user.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   fmt.Sprint(err),
			"data":   nil,
		})
		return
	}

	//生成access_token和refresh_token
	resp, err := rpc.KeyCenter.CreateToken(context.Background(), &rpc.CreateTokenReq{
		Domain: "common",
		Uid:    u.Uid,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"status": 50000,
			"info":   "err",
			"data":   nil,
		})
		return
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
