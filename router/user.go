package router

import (
	"gintest/types/model"
	"gintest/utils"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespErr(c, "解析出错")
		return
	}

	_, err := model.CreateUser(user)
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	utils.RespSuc(c)
	return
}
