package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"tool_update_service/logic"
	"tool_update_service/models"
)

func ReleaseHandler(c *gin.Context) {
	ret, err := logic.Release()
	if err != nil {
		zap.L().Info("暂无版本号！")
	}
	//c.HTML(http.StatusOK, "release/upload.html", ret)
	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"msg":     "获取最新版本成功",
		"version": ret,
	})
}

func UpdateHandler(c *gin.Context) {
	//1.获取参数和参数校验
	//ShouldBindJSON只能校验字段类型对不对，请求数据的格式，譬如是不是json格式
	var p = new(models.ParamUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("update param invalid.", zap.Error(err))
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
	info, err := logic.GetUpdateinfo(p)
	if err != nil {
		zap.L().Error("无需更新...", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": 10009,
			"msg":  "无需更新",
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":     "可以更新啦~",
		"code":    10010,
		"version": info.Version,
		"content": info.Content,
		"forced":  info.Forced,
		"url":     info.Url,
		"size":    info.Size,
	})
	zap.L().Info("获取更新信息成功....")
}

// GetUpdateHistoryHandler 获取历史更新记录，进行参数校验和业务处理转发及返回响应
func GetUpdateHistoryHandler(c *gin.Context) {

	//c.HTML(http.StatusOK, "download/download.html", nil)
	/*tmpl, err := template.ParseFiles("C:\\Users\\dell\\go\\src\\tool_update_service\\templates\\download\\download.html")
	if err != nil {
		zap.L().Error("tmpl error.", zap.Error(err))
	}*/
	// 转发至logic层进行业务处理

	info, err := logic.GetUpdateHistory()
	if err != nil {
		zap.L().Error("获取更新历史记录失败")
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "获取更新历史记录失败!",
		})
		return
	}

	// 返回响应结果
	fmt.Printf("info:%v\n", info[0])
	for _, r := range info {
		fmt.Printf("%v\n", r)
	}
	/*
		if err := tmpl.Execute(c.Writer, info); err != nil {
			zap.L().Error("写入模板失败！")
		}*/

	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "获取历史更新记录成功",
		"data": info,
	})

	zap.L().Info("获取历史更新记录成功")
}
