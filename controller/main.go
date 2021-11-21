package controller

import (
	"github.com/gin-gonic/gin"
)

func ReturnFalse(c *gin.Context) {
	c.JSON(406, gin.H{
		"code":    406,
		"message": "参数错误",
	})
}
func ReturnTrue(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
