package c

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"server/d"
	"server/pkg/common"
	"server/pkg/e"
)

func Post(c *gin.Context) {
	var task d.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		common.CJSON(c, http.StatusOK, e.INVALID_PARAMS)
		return
	}

	_, err = d.AddTask(c.Keys["uid"].(int64), &task)
	if err != nil {
		if _, ok := err.(*mysql.MySQLError); ok {
			common.CJSON(c, http.StatusOK, e.ERROR_POST_M)
			return
		}
		common.CJSON(c, http.StatusOK, e.ERROR_POST_C)
		return
	}

	common.CJSON(c, http.StatusOK, e.SUCCESS)
	return
}

func GetTaskList(c *gin.Context) {
	return
}
