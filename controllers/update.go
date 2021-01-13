package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"tool_update_service/dao/mysql"
	"tool_update_service/logic"
	"tool_update_service/models"
)

func UpdateHandler(c *gin.Context) {
	//1.获取参数和参数校验
	//ShouldBindJSON只能校验字段类型对不对，请求数据的格式，譬如是不是json格式
	var p = new(models.ParamUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
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

	//2.业务处理
	if err := logic.Update(p); err != nil {
		zap.L().Error("发布失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "发布失败",
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "发布成功",
	})
	zap.L().Info("发布成功...")
}

// GetUpdateinfoHandler 对获取更新信息请求进行参数校验和业务处理转发
func GetUpdateinfoHandler(c *gin.Context) {
	versionNum := c.Query("version")
	if len(versionNum) <= 0 {
		zap.L().Error("getupdateinfo param invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": 40004,
			"msg":  "参数有误",
		})
		return
	}

	//2.业务处理
	p := new(models.ParamGetUpdateinfo)
	p.Version = versionNum
	if err := logic.GetUpdateinfo(p); err != nil {
		zap.L().Error("无需更新...", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 10009,
			"msg":  "无需更新",
		})
		return
	}
	//3.返回响应
	info := mysql.Info
	c.JSON(http.StatusOK, gin.H{
		"msg":     "可以更新啦~",
		"code":    10010,
		"version": info.Version,
		"content": info.Content,
		"url":     info.Url,
	})
	zap.L().Info("获取更新信息成功....")
}

func GetUpdateListHandler(c *gin.Context) {

}
