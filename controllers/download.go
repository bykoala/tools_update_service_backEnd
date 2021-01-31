package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"tool_update_service/logic"
)

func DownloadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "download.html", nil)
	filePath := c.Request.RequestURI
	zap.L().Info(filePath)
	//  查询一些必要的参数 进行一些必要的验证
	//attchIdStr := c.Query("attachment_id")
	//attachmentName := c.Query("attachment_name")
	//filePath := c.Query("filepath")
	all, err := logic.Download(filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "下载文件失败,请联系苏宝伢",
		})
	}
	// 设置返回头并返回数据
	//fileContentDisposition := "attachment;filename=\"" + attachmentName + "\""
	c.Header("Content-Type", "java/*") // 这里是压缩文件类型 .zip
	dfn := strings.Split(filePath, "/")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", dfn[len(dfn)-1]))
	//c.Header("Content-Disposition", fileContentDisposition)
	c.Data(http.StatusOK, "application/zip", all)
}
