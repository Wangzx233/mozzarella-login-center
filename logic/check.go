package logic

import (
	"errors"
	"log"
	"mozzarella-login-center/model"
	"regexp"
)

// Check 检查用户的电话号，密码，用户类型的格式
func Check(u *model.UserJson) (err error) {

	err = CheckPhoneNumber(u.PhoneNumber)

	return
}

// CheckPhoneNumber 正则表达式检查电话号格式
func CheckPhoneNumber(phoneNumber string) (err error) {
	compilePhoneNumber := regexp.MustCompile("^1[3,4,5,7,8][0-9]{9}$")
	if compilePhoneNumber == nil {
		log.Println("regexp err")
		return errors.New("err")
	}
	res := compilePhoneNumber.FindStringSubmatch(phoneNumber)
	if res == nil {

		return errors.New("电话号格式错误")
	}
	return nil
}

// CheckUserType 验证用户类型格式
func CheckUserType(userType int) (err error) {
	if userType < 0 || userType > 2 {
		err = errors.New("用户类型错误")
	}
	return
}

// CheckPassword 验证密码：密码由6-16位字符串组成，必须包含数字、字母、符号中至少两种元素
func CheckPassword(password string) (err error) {
	err = errors.New("密码由6-16位字符串组成，必须包含数字、字母、符号中至少两种元素")

	var element [3]bool //0:数字 1：字母 2：符号
	l := len(password)
	if l < 6 || l > 16 {
		return
	}

	for _, c := range password {

		if c >= '0' && c <= '9' {
			element[0] = true
		} else {
			if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
				element[1] = true
			} else {
				if c == '*' || c == '=' || c == '_' || c == '-' {
					element[2] = true
				} else {
					return errors.New("密码只支持数字、字母、*、+、-、_")
				}
			}
		}

	}

	//检查是否至少包含两种元素
	if element[0] && element[1] || element[1] && element[2] || element[0] && element[2] {
		return nil
	}

	return errors.New("密码格式错误")
}
