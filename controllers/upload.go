package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"tool_update_service/logic"
	"tool_update_service/models"
	"tool_update_service/tools/upload"
)

func UploadHandler(c *gin.Context) {
	//参数校验

	var p = new(models.ParamUpload)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("update param invalid", zap.Error(err))
		//判断err是不是validator校验的类型
		e, ok := err.(validator.ValidationErrors)
		//如果不是，则不翻译
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 40004,
				"msg":  "参数有误",
			})
			return
		}

		//如果是，则翻译
		c.JSON(http.StatusOK, gin.H{
			"code": 40001,
			"msg":  removeTopStruct(e.Translate(trans)),
		})
		return

	}
	//逻辑处理转发
	if err := logic.Upload(p); err != nil {
		if err.Error() == "fileIsExist" {
			zap.L().Error("文件已存在,请勿重复上传")
			//根据文件md5值，判断是否已存在文件
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "文件已存在,请勿重复上传",
			})
			return
		} else {
			zap.L().Error("文件上传异常")
			c.JSON(http.StatusOK, gin.H{
				"code": 10004,
				"msg":  "文件上传异常",
			})
			return
		}
	}

	if err := upload.File(c, p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10009,
			"msg":  "文件文件上传失败,可能文件为空",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "文件上传成功",
		"url":  p.Url,
		"size": p.Size,
		"md5":  p.Md5,
	})

}
