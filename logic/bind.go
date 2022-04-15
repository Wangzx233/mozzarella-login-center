package logic

import (
	uuid "github.com/satori/go.uuid"
	"mozzarella-login-center/dao"
	"mozzarella-login-center/model"
)

func Bind(domain string, openid string, union, studentID string, realName string) (err error) {
	var user = model.User{
		Uid:         uuid.NewV4().String(),
		StudentID:   studentID,
		RealName:    realName,
		PhoneNumber: "",
		UnionID:     union,
	}
	switch domain {
	case "xcx":
		user.XcxOpenID = openid
	case "app":
		user.AppOpenID = openid
	}

	err = dao.Register(user)
	return
}
