package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mozzarella-login-center/dao"
	"mozzarella-login-center/logic"
	"mozzarella-login-center/model"
	"mozzarella-login-center/resps"
	"mozzarella-login-center/rpc"
	"strconv"
)

var (
	appid           = "wxc58c92b1f02c1933"
	appSecret       = "bb33746ef6e760afaa2f928cc1cddc02"
	code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=" + appid + "&secret=" + appSecret + "&js_code=%s&grant_type=authorization_code"
)

// XcxLogin copy by rashB
func XcxLogin(c *gin.Context) {
	//1.得到code
	code := c.Query("code")
	//2.发送授权得到openid
	if code == "" {
		log.Println("code为空")
		resps.HandleError(c, errors.New("code为空"))
		return
	}

	session, err := getJsCodeSession(code)
	if err != nil {
		log.Println("微信小程序登录凭证校验失败")
		resps.HandleError(c, err)
		return
	}

	stuNum := c.PostForm("student_id")
	realName := c.PostForm("real_name")

	//3.判断此openid是否绑定用户数据
	user, err := dao.FindUserByOpenid("xcx", session.Openid)
	//若得到stuNum和realName说明需要绑定绑定
	if stuNum != "" && realName != "" && user.StudentID != stuNum {
		//检查学号和姓名正确性
		ok, err := dao.VerifyUser(stuNum, realName)
		if !ok {
			resps.ParamError(c)
			return
		}
		//根据stuNum查询是否绑定其他应用
		err = logic.Bind("xcx", session.Openid, session.UnionId, stuNum, realName)
		if err != nil {
			log.Println("error in Bind", err)
			resps.HandleError(c, err)
			return
		}
	}

	//3.1若没有绑定则返回提示信息，由前端跳转到绑定页面附带stuNum和realName再次请求该接口
	if err != nil {
		log.Println("openid[" + session.Openid + "] 需要绑定信息")
		resps.ParamError(c)
		return
	}

	resp, err := rpc.KeyCenter.CreateToken(context.Background(), &rpc.CreateTokenReq{
		Domain: "xcx",
		Uid:    user.Uid,
	})

	resps.OKWithData(c, gin.H{
		"access_token":  resp.Token,
		"refresh_token": resp.RefreshToken,
	})
}

//getJsCodeSession 调用code2Session获取JsCode2Session返回信息
func getJsCodeSession(code string) (session model.JsCodeSessionResponse, err error) {
	res, _ := resps.SendGet(fmt.Sprintf(code2SessionURL, code))

	err = json.Unmarshal(res, &session)

	if err != nil || session.ErrCode != 0 {
		err = errors.New("get code2Session error(" + strconv.Itoa(session.ErrCode) + "):" + session.ErrMsg)
		return
	}

	return
}
