package logic

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"mozzarella-login-center/dao"
	"mozzarella-login-center/model"
)

func Bind(domain string, openid string, union, studentID string, realName string) (err error) {
	isRegister, u := dao.IsRegisterStudentID(studentID)
	var user = model.User{
		StudentID: studentID,
		RealName:  realName,
		UnionID:   union,
	}

	//判断对应平台是否已经注册
	switch domain {
	case "xcx":
		if isRegister {
			//如果已经注册且对应openid不为空说名在本平台已注册
			if u.XcxOpenID != "" {
				return errors.New("用户已注册")
			}
			user.Uid = u.Uid
		}
		user.XcxOpenID = openid
	case "app":
		if isRegister {
			if u.AppOpenID != "" {
				return errors.New("用户已注册")
			}
			user.Uid = u.Uid
		}
		user.AppOpenID = openid
	}

	if isRegister {
		err = dao.UpdateUser(user)
	} else {
		user.Uid = uuid.NewV4().String()
		err = dao.Register(user)
	}
	return
}
