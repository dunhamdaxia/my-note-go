package router

import (
	"fmt"
	"gintest/utils"
	"testing"
)

var (
	createTaskUrl   = "task/create"
	pageTaskUrl     = "task/page"
	completeTaskUrl = "task/complete"
	token           = "$2a$04$lOCpjWr/EA0tS3W3hxtJcOm7X/bbM08D9nXj1iQX18HWuCGXF7Cm2"
)

func TestCreateTask(t *testing.T) {
	postData := utils.M{
		"name":        "This is test",
		"period_type": 2,
	}

	res, code, err := utils.HttpPostJson(host+createTaskUrl, postData, utils.M{"token": token})
	fmt.Println("statusCode", code)
	fmt.Println("response:", string(res))
	fmt.Println("err:", err)
}

func TestPageTask(t *testing.T) {
	postData := utils.M{
		"page":   1,
		"num":    10,
		"status": 1,
	}

	res, code, err := utils.HttpPostJson(host+pageTaskUrl, postData, utils.M{"token": token})
	fmt.Println("statusCode", code)
	fmt.Println("response:", string(res))
	fmt.Println("err:", err)
}

func TestCompleteTask(t *testing.T) {
	postData := utils.M{
		"id": 1,
	}

	res, code, err := utils.HttpPostJson(host+completeTaskUrl, postData, utils.M{"token": token})
	fmt.Println("statusCode", code)
	fmt.Println("response:", string(res))
	fmt.Println("err:", err)
}
