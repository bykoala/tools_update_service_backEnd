package routers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"tool_update_service/controllers"
	"tool_update_service/logger"
	"tool_update_service/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middleware.CrosMiddleware()) //允许跨域

	//gin框架给模板添加自定义方法
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.Static("xxx", "./statics/")
	//指定模板路径
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/debugajax.tmpl")
	//发版页面接口
	r.GET("/release", controllers.ReleaseHandler)
	//发版信息接口
	r.POST("/update", controllers.UpdateHandler)
	// GetUpdateinfo 该接口提供给客户端查询是否需要更新，客户端需要送参version
	//根据version从数据库查找当前最新一条记录，如果有，则无需更新；如果无，则需要更新。
	//"code": 10009,"msg":  "无需更新"
	//"code": 10010,"msg":  "需要更新"
	//获取更新信息接口
	r.GET("/getUpdateinfo", controllers.GetUpdateinfoHandler)
	//历史更新记录接口
	r.GET("/getUpdateHistory", controllers.GetUpdateHistoryHandler)
	//文件上传接口
	r.POST("/upload", controllers.UploadHandler)
	//文件下载接口
	r.GET("/download/*path", controllers.DownloadHandler)
	//展示反馈或建议的页面
	r.GET("/getFeedbackOrBug", controllers.GetFeedbackOrBud)
	r.POST("/feedbackOrBug", controllers.FeedbackOrBugHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "404",
		})
	})

	return r
}
