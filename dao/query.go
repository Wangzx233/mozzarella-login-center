package dao

import (
	"gorm.io/gorm"
	"mozzarella-login-center/model"
)

// FindUserByPhoneNumber 通过手机号查询用户
func FindUserByPhoneNumber(phone string) (user model.User, err error) {
	err = db.Model(&model.User{}).Where("phone_number = ?", phone).First(&user).Error
	return
}

// VerifyUser 验证学号和姓名是否存在匹配
func VerifyUser(studentID, realName string) (ok bool, err error) {
	err = db.Model(&model.Student{}).Where("student_id = ? and real_name =?", studentID, realName).First(&model.Student{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ok = false
			err = nil
		}
		return
	}
	ok = true
	return
}

// FindUserByUid uid找到用户
func FindUserByUid(uid string) (user model.User, err error) {
	err = db.Model(&model.User{}).Where("uid = ?", uid).First(&user).Error
	return
}

// FindUserByOpenid openid找到用户
func FindUserByOpenid(domain, openid string) (user model.User, err error) {
	switch domain {
	case "xcx":
		err = db.Model(&model.User{}).Where("xcx_openid = ?", openid).First(&user).Error
	case "app":
		err = db.Model(&model.User{}).Where("app_openid = ?", openid).First(&user).Error
	}

	return
}
