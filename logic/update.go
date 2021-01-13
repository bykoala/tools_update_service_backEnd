package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"tool_update_service/dao/mysql"
	"tool_update_service/models"
)

func Update(p *models.ParamUpdate) (err error) {
	//1.先判断版本是否已存在
	if err := mysql.VersionIsExist(p.Version); err != nil {
		zap.L().Error("发布失败")
		return err
	}

	//2.构造update实例
	u := &models.UpdateInfo{
		Version: p.Version,
		Content: p.Content,
		Forced:  p.Forced,
		Url:     p.Url,
		Status:  p.Status,
	}

	//3.写入版本更新信息
	if err := mysql.InsertUpdateInfo(u); err != nil {
		return err
	}
	fmt.Printf("更新成功!")
	return
}

func GetUpdateinfo(p *models.ParamGetUpdateinfo) (err error) {
	if err := mysql.QueryLatestRecord(); err != nil {
		zap.L().Error("数据库没有查询到数据")
		return err
	}
	info := mysql.Info
	//客户端送来的版本跟数据库查找出来的版本号进行比对
	fmt.Printf("p.verson:%v,info.version:%v\n", p.Version, info.Version)
	if p.Version == info.Version {
		zap.L().Info("无需更新")
		return errors.New("无需更新")
	}
	//fmt.Printf("version:%v,content:%v,forced:%v,status:%v,url:%v,create_time:%v\n", info.Version, info.Content, info.Forced, info.Status, info.Url, info.CreateTiem)

	return
}
