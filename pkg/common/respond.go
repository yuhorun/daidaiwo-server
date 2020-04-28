package common

import (
	"github.com/gin-gonic/gin"
	"server/pkg/e"
)

func CJSON(c *gin.Context, status int, code int) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
