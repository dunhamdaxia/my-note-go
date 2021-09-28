package router

import (
	"gintest/stores/mysql"
	"gintest/types/model"
	"gintest/types/param"
	"gintest/utils"
	"github.com/gin-gonic/gin"
)

func createTask(c *gin.Context) {
	token := c.GetHeader("token")

	user, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	taskData := param.TaskCreate{}
	if err := c.ShouldBindJSON(&taskData); err != nil {
		utils.RespErr(c, "解析数据出错")
		return
	}

	if taskData.Name == "" {
		utils.RespErr(c, "缺少必要的参数:Name")
		return
	}

	if taskData.PeriodType == 0 {
		utils.RespErr(c, "缺少必要的参数:PeriodType")
		return
	}

	newTaskData := model.Task{
		Name:       taskData.Name,
		PeriodType: taskData.PeriodType,
		UserID:     user.ID,
		Desc:       taskData.Desc,
	}

	_, err := model.CreateTask(newTaskData)
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	utils.RespSuc(c)
}

func pageTask(c *gin.Context) {
	token := c.GetHeader("token")

	user, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	paramData := param.TaskPage{}
	if err := c.ShouldBindJSON(&paramData); err != nil {
		utils.RespErr(c, "解析出错")
		return
	}

	db, err := mysql.NewDb()
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	where := map[string]interface{}{"user_id": user.ID}
	if paramData.Status > 0 {
		where["status"] = paramData.Status
	}

	tasks := make([]model.Task, 0)
	db.Where(where).Limit(paramData.Num).Offset(paramData.Num * (paramData.Page - 1)).Find(&tasks)

	if db.Error != nil {
		utils.RespErr(c, db.Error.Error())
		return
	}

	utils.RespSuc(c, "ok", tasks)
}

func completeTask(c *gin.Context) {
	token := c.GetHeader("token")

	user, code, msg := model.CheckToken(token)
	if code != 2 {
		utils.RespErr(c, msg, utils.M{"code": code})
		return
	}

	paramData := struct {
		ID int64
	}{}
	if err := c.ShouldBindJSON(&paramData); err != nil {
		utils.RespErr(c, "解析出错")
		return
	}

	db, err := mysql.NewDb()
	if err != nil {
		utils.RespErr(c, err.Error())
		return
	}

	task := model.Task{}
	db.Where(map[string]interface{}{"id": paramData.ID, "user_id": user.ID}).First(&task)

	if task == (model.Task{}) {
		utils.RespErr(c, "未找到相关任务")
		return
	}

	if task.Status != 1 {
		utils.RespErr(c, "任务状态不支持修改")
		return
	}

	//todo 开始事务
	db.Model(&task).Update("status", 2)
	if db.Error != nil {
		utils.RespErr(c, db.Error.Error(), nil)
		return
	}

	//todo 奖励发放

	utils.RespSuc(c, "ok")
}
