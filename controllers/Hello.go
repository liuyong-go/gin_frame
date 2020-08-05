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
	// time.Sleep(10 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功1",
		"data": result,
	})
	return
}
func TestHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功2",
		"data": "result",
	})
	return
}
