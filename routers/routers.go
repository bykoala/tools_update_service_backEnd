package routers

import (
	"github.com/gin-gonic/gin"
	"tool_update_service/controllers"
	"tool_update_service/logger"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/update", controllers.UpdateHandler)
	// GetUpdateinfo 该接口提供给客户端查询是否需要更新，客户端需要送参version
	//根据version从数据库查找当前最新一条记录，如果有，则无需更新；如果无，则需要更新。
	//"code": 10009,"msg":  "无需更新"
	//"code": 10010,"msg":  "需要更新"
	r.GET("/getUpdateinfo", controllers.GetUpdateinfoHandler)
	r.GET("/getUpdateList", controllers.GetUpdateListHandler)

	/*	r.GET("/getUpdateinfo", func(c *gin.Context) {
		//time.Sleep(time.Second*10)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": "pong",
		})

	})*/
	return r

}
