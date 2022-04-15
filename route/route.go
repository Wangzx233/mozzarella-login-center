package route

import (
	"github.com/gin-gonic/gin"
	"mozzarella-login-center/api"
)

func InitRoute() {
	engine := gin.Default()

	user := engine.Group("/login-center")

	{
		xcx := user.Group("/xcx")
		{
			xcx.POST("/login", api.XcxLogin)
		}

		user.POST("/login", api.LoginByCode)
		user.POST("/register", api.Register)
		user.GET("/check", api.Check)

		user.GET("/login", api.SendLoginCode) //发送登陆短信验证码
		user.GET("/register", api.SendRegisterCode)
	}

	engine.Run(":8085")
}
