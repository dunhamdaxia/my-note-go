package router

import (
	"fmt"
	"gintest/utils"
	"testing"
)

var (
	host          = "http://127.0.0.1:8080/"
	createUserUrl = "create/user"
)

func TestCreatUser(t *testing.T) {
	data := utils.M{
		"name":     "harden",
		"account":  "account2",
		"password": "123456",
	}
	res, code, err := utils.HttpPostJson(host+createUserUrl, data)

	fmt.Println("statusCode", code)
	fmt.Println("response:", string(res))
	fmt.Println(err)
}
