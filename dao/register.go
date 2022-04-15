package dao

import (
	"gorm.io/gorm"
	"log"
	"mozzarella-login-center/model"
)

func Register(u model.User) (err error) {
	err = db.Create(&u).Error
	return
}

// IsRegisterPhoneNumber 判断用户是否已存在，true为存在
func IsRegisterPhoneNumber(phoneNumber string) (IsRegister bool) {
	IsRegister = true
	err := db.Model(&model.User{}).Where("phone_number = ?", phoneNumber).First(model.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			IsRegister = false
		} else {
			log.Println(err)
		}
	}
	return
}

// IsRegisterStudentID 判断StudentID是否已存在，true为存在
func IsRegisterStudentID(StudentID string) (IsRegister bool) {
	IsRegister = true
	err := db.Model(&model.User{}).Where("student_id = ?", StudentID).First(model.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			IsRegister = false
		} else {
			log.Println(err)
		}
	}
	return
}
