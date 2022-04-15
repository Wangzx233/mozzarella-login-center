package logic

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"math/rand"
	"mozzarella-login-center/dao"
	"strconv"
	"time"
)

func SendCode(phoneNumber string) (code string, err error) {
	client, err := CreateClient(tea.String("LTAI5tMgY1RqYU1fikp5feaH"), tea.String("TDtSUGQM1QmTt97Oxks5rEe45ef7qG"))
	if err != nil {

		return
	}
	code = GenerateCode()
	err = dao.SaveCode(phoneNumber, code)
	if err != nil {
		return
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	// 复制代码运行请自行打印 API 的返回值
	resp, err := client.SendSms(sendSmsRequest)
	if err != nil {
		return
	}
	log.Println(resp)
	return
}
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(899999) + 100000

	code := strconv.Itoa(number)
	return code
}
