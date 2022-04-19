package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-login-center/dao"
	"mozzarella-login-center/resps"
)

func IsRegister(c *gin.Context) {
	studentId := c.Query("student_id")
	domain := c.Query("domain")

	is, user := dao.IsRegisterStudentID(studentId)
	if is {
		switch domain {
		case "xcx":
			if user.XcxOpenID == "" {
				is = false
			}
		case "app":
			if user.AppOpenID == "" {
				is = false
			}
		}
	}

	resps.OKWithData(c, gin.H{
		"is_register": is,
	})
}

// Check 检查学号和姓名是否合法
func Check(c *gin.Context) {
	studentID := c.Query("student_id")
	realName := c.Query("real_name")

	ok, err := dao.VerifyUser(studentID, realName)
	if err != nil {
		log.Println("err:", err)
		c.JSON(503, gin.H{
			"status": 50000,
			"info":   " err ",
		})
		return
	}

	if !ok {
		c.JSON(401, gin.H{
			"status": 40000,
			"info":   "学生不存在",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 10000,
		"info":   "success",
	})
}
