package router

import (
	"fmt"
	"gintest/utils"
	"testing"
)

var (
	loginUrl      = "login"
	checkTokenUrl = "check_token"
)

func TestLogin(t *testing.T) {
	data := utils.M{
		"account":  "account1",
		"password": "123456",
	}

	res, code, err := utils.HttpPostJson(host+loginUrl, data)
	fmt.Println("statusCode", code)
	fmt.Println("response:", string(res))
	fmt.Println("err:", err)
}

func TestCheckToken(t *testing.T) {
	res, code, err := utils.HttpPostJson(host+checkTokenUrl, utils.M{})
	fmt.Println("statusCode", code)
	fmt.Println("response:", string(res))
	fmt.Println("err:", err)
}
