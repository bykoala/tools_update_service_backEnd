package logic

import (
	"errors"
	"go.uber.org/zap"
	"tool_update_service/dao/mysql"
	"tool_update_service/models"
	"tool_update_service/tools/fileInfo"
)

/*func Upload(p *models.ParamUpload) (err error) {
	//tools.CalcMd5()

	err = mysql.VersionIsExist(p.Version)
	//客户端送来的版本跟数据库查找出来的版本号进行比对
	if err != nil {
		if err.Error() == "versionIsExist" {
			zap.L().Info("已存在一样的版本")
			return errors.New("fileIsExist")
		}
	}
	return nil
}
*/
func Upload(p *models.ParamUpload) (err error) {
	md5, err := fileInfo.CalcMd5(p.FileName)
	if err != nil {
		zap.L().Error("获取文件的md5值失败！", zap.Error(err))
		return errors.New("getFileMd5Failed")
	}
	p.Md5 = md5
	err = mysql.Md5IsExist(md5)
	//客户端送来的版本跟数据库查找出来的版本号进行比对
	if err != nil {
		if err.Error() == "md5IsExist" {
			zap.L().Info("已存在一样的文件")
			return errors.New("fileIsExist")
		} else {
			panic("未知错误")
		}
	}

	return nil
}
