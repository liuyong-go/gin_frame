package controllers

import (
	"gin_frame/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")
	maps := make(map[string]interface{})
	maps["username"] = mobile
	result := models.GetAdmin(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": result,
	})
	return
}
