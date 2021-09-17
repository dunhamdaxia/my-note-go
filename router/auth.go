package router

import (
	"gintest/types/model"
	"gintest/types/param"
	"gintest/utils"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	paramData := param.Login{}
	if err := c.ShouldBindJSON(&paramData); err != nil {
		utils.RespErr(c, "解析出错")
		return
	}

	token, err := model.Login(paramData.Account, paramData.Password)
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	utils.RespSuc(c, "ok", utils.M{"token": token})
	return
}

func checkToken(c *gin.Context) {
	token := c.GetHeader("token")

	user, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	utils.RespSuc(c, msg, utils.M{"user": user, "code": code})
}
