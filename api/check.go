package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-login-center/dao"
)

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
