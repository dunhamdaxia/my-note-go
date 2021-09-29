package router

import (
	"gintest/stores/mysql"
	"gintest/types/model"
	"gintest/types/param"
	"gintest/utils"
	"github.com/gin-gonic/gin"
)

func saveAttribute(c *gin.Context) {
	token := c.GetHeader("token")

	user, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	attributeData := param.AttributeCreate{}
	if err := c.ShouldBindJSON(&attributeData); err != nil {
		utils.RespErr(c, "解析数据出错")
		return
	}

	if attributeData.Name == "" {
		utils.RespErr(c, "缺少必要的参数:Name")
		return
	}

	newAttributeData := model.Attributes{
		Name:     attributeData.Name,
		ParentID: attributeData.ParentID,
		Sort:     attributeData.Sort,
		UserID:   user.ID,
	}

	var err error
	if attributeData.ID == 0 {
		_, err = model.CreateAttribute(newAttributeData)
	} else {
		err = model.UpdateAttribute(newAttributeData)
	}

	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	utils.RespSuc(c, "ok")
}

func pageAttributes(c *gin.Context) {
	token := c.GetHeader("token")

	user, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	paramData := param.AttributesPage{}
	if err := c.ShouldBindJSON(&paramData); err != nil {
		utils.RespErr(c, "解析出错")
		return
	}

	db, err := mysql.NewDb()
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	where := map[string]interface{}{"user_id": user.ID, "parent_id": paramData.ParentId}

	attributes := make([]model.Attributes, 0)
	db.Order("sort desc").Where(where).Limit(paramData.Num).Offset(paramData.Num * (paramData.Page - 1)).Find(&attributes)

	if db.Error != nil {
		utils.RespErr(c, db.Error.Error())
		return
	}

	utils.RespSuc(c, "ok", attributes)
}

func infoAttribute(c *gin.Context) {
	token := c.GetHeader("token")

	_, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	paramData := param.AttributeInfo{}
	if err := c.ShouldBindJSON(&paramData); err != nil {
		utils.RespErr(c, "解析出错")
		return
	}

	if paramData.ID == 0 {
		utils.RespErr(c, "缺少必要的参数:id")
		return
	}

	db, err := mysql.NewDb()
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	attribute := model.Attributes{}
	db.Where(map[string]interface{}{"id": paramData.ID}).First(&attribute)

	if db.Error != nil {
		utils.RespErr(c, err.Error())
		return
	}

	utils.RespSuc(c, "ok", attribute)
}
