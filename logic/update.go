package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"tool_update_service/dao/mysql"
	"tool_update_service/models"
)

// Release 查找数据库最近一条数据，并返回
func Release() (v string, err error) {

	ret, err := mysql.QueryLatestRecord()

	if err != nil {
		zap.L().Error("数据查询失败")
		if err.Error() == "notResult" {
			return
		} else {
			panic("kda")
		}
	}
	return ret.Version, nil
}

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
		Size:    p.Size,
		Md5:     p.MD5,
		Status:  p.Status,
	}

	//3.写入版本更新信息
	if err := mysql.InsertUpdateInfo(u); err != nil {
		return err
	}
	fmt.Printf("更新成功!")
	return
}

func GetUpdateinfo(p *models.ParamGetUpdateinfo) (info *models.UpdateInfo, err error) {
	info, err = mysql.QueryLatestRecord()
	if err != nil {
		zap.L().Error("数据库没有查询到数据")
		return nil, err
	}
	//客户端送来的版本跟数据库查找出来的版本号进行比对
	if p.Version == info.Version {
		zap.L().Info("无需更新")
		return nil, errors.New("无需更新")
	}
	return
}

func GetUpdateHistory() (info []*models.UpdateInfo, err error) {
	return mysql.QueryUpdateHistory()
}
