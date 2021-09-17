package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErr(c *gin.Context, data ...interface{}) {
	returnData := M{
		"status": false,
		"msg":    "error",
	}

	dataLen := len(data)
	if dataLen > 0 {
		returnData["msg"] = data[0]
	}

	if dataLen > 1 {
		returnData["data"] = data[1]
	}

	c.JSON(http.StatusOK, returnData)
}

func RespSuc(c *gin.Context, data ...interface{}) {
	returnData := M{
		"status": false,
		"msg":    "ok",
	}

	dataLen := len(data)
	if dataLen > 0 {
		returnData["msg"] = data[0]
	}

	if dataLen > 1 {
		returnData["data"] = data[1]
	}

	c.JSON(http.StatusOK, returnData)
}
