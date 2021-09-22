package router

import (
	"gintest/router/routercommon"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(routercommon.Cors())

	r.GET("/topgoer", helloHandler)
	r.POST("/register", createUser)
	r.POST("/login", login)
	r.POST("/task/create", createTask)
	r.POST("/task/page", pageTask)
	r.POST("/task/complete", completeTask)
	r.POST("/check_token", checkToken)
	return r
}
