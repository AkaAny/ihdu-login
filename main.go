package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

const (
	USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
)

func main() {
	loginInfo, err := GetLoginInfo()
	if err != nil {
		panic(err)
	}
	var url = "http://2.2.2.2/ac_portal/login.php"
	var request = gorequest.New()
	request.DoNotClearSuperAgent = true
	request.Header.Add("User-Agent", USER_AGENT)
	request.Post(url).Type("form")
	resp, body, errs := request.Send(map[string]interface{}{
		"opr":         "pwdLogin",
		"userName":    loginInfo.UserName,
		"pwd":         loginInfo.Password,
		"rememberPwd": 0,
	}).End()
	if errs != nil {
		panic(errs[0])
	}
	fmt.Println(resp.StatusCode)
	fmt.Printf("response:\n%s", body)
	err = ioutil.WriteFile("resp.json", []byte(body), 0644)
	if err != nil {
		panic(err)
	}
}
