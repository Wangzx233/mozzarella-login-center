package logic

import (
	uuid "github.com/satori/go.uuid"
	"mozzarella-login-center/model"
)

// UserJsonToDao 将user_json转换为user_dao
func UserJsonToDao(json model.UserJson) (user model.User) {
	user = model.User{
		Uid:         uuid.NewV4().String(),
		StudentID:   json.StudentID,
		RealName:    json.RealName,
		PhoneNumber: json.PhoneNumber,
	}
	return
}
