package logic

import (
	"errors"
	"log"
	"mozzarella-login-center/dao"
	"mozzarella-login-center/model"
)

// LoginByCode 验证码登陆
func LoginByCode(u *model.UserJson) (err error) {
	code, err := dao.GetCode(u.PhoneNumber)
	if err != nil {
		log.Println("get code err : ", err)
		err = errors.New("验证码错误")
		return
	}

	if code != u.Code {
		err = errors.New("验证码错误")
	}

	return
}
