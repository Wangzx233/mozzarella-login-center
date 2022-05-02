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

		user.GET("/isRegister", api.IsRegister) //判断用户是否已经注册

		user.POST("/login", api.LoginByCode)
		user.POST("/register", api.Register)
		user.GET("/check", api.Check)

		user.GET("/login", api.SendLoginCode) //发送登陆短信验证码
		user.GET("/register", api.SendRegisterCode)

		user.GET("/refresh", api.Refresh)

		app := user.Group("/app")
		{
			app.GET("/getTestToken", api.GetTestToken) //测试临时接口
		}
	}

	engine.Run(":8085")
}
