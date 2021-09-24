package router

import (
	"gintest/types/model"
	"gintest/types/param"
	"gintest/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func createUser(c *gin.Context) {
	paramsData := param.Register{}
	if err := c.ShouldBindJSON(&paramsData); err != nil {
		utils.RespErr(c, "解析出错|"+err.Error())
		return
	}

	birthdayTime, err := time.ParseInLocation("2006-01-02", paramsData.Birthday, time.Local)
	if err != nil {
		utils.RespErr(c, "出生日期转换失败|"+err.Error())
		return
	}
	user := model.User{
		Phone:    paramsData.Phone,
		Name:     paramsData.Name,
		Height:   paramsData.Height,
		Weight:   paramsData.Weight,
		Account:  paramsData.Account,
		Password: paramsData.Password,
		Birthday: birthdayTime,
	}
	_, err = model.CreateUser(user)
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	utils.RespSuc(c)
	return
}
