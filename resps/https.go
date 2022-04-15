package resps

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var (
	c *http.Client
)

func SendGet(url string) (b []byte, err error) {
	body, err := c.Get(url)
	if err != nil {
		return
	}
	b, err = ioutil.ReadAll(body.Body)
	return
}

func SendPost(url string, data map[string]string) (b []byte) {
	param := ""
	for k, v := range data {
		param += k + "=" + v + "&"
	}
	param = param[:len(param)-1]
	r := bytes.NewReader([]byte(param))
	resp, err := c.Post(url, "application/x-www-form-urlencoded", r)
	if err != nil {
		log.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return body
}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)

	c = &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
}
