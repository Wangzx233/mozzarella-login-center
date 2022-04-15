package logic

import "mozzarella-login-center/dao"

func IsRegister(phoneNumber, studentId string) (IsRegisterP bool, IsRegisterS bool) {
	IsRegisterS = dao.IsRegisterStudentID(studentId)
	IsRegisterP = dao.IsRegisterPhoneNumber(phoneNumber)
	return
}
