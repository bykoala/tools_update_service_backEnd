package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path"
	"strings"
	"tool_update_service/dao/mysql"
	"tool_update_service/models"
	"tool_update_service/tools"
	"tool_update_service/tools/fileInfo"
)

// Release 查找数据库最近一条数据，并返回
func Release() (v string, err error) {

	ret, err := mysql.QueryLatestRecord()

	if err != nil {
		zap.L().Error("数据查询失败")
		if err.Error() == "notResult" {

			return
		} else {
			panic("kendiea")
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
		Version:        p.Version,
		Content:        p.Content,
		Forced:         p.Forced,
		Url:            p.Url,
		Size:           p.Size,
		Md5:            p.MD5,
		Status:         p.Status,
		Classification: p.Classification,
	}
	//3.对上传文件逻辑改造后，增加文件归类存放

	dst := u.Url // http://192.168.3.136:8088/download/files/??????.exe
	fileUrls := strings.Split(dst, "files")
	url_base := fileUrls[0] //http://192.168.3.136:8088/download/
	fileName := fileUrls[1]
	newDir := ""
	dst = path.Join(url_base, "./files/img", fileUrls[1])
	fmt.Printf("classification:%v\n", p.Classification)
	if p.Classification == 0 {
		//拼接成：http://192.168.3.136:8088/download/files/source/v1.0.0/??????.exe
		dst = fmt.Sprintf("%s%s%s", url_base, "./files/img", fileUrls[1])
	} else {
		//准备存放文件的目录：
		newDir = "./files/source/" + p.Version
		tools.Mkdir(newDir)
		//起一个协程复制文件
		go func() {
			//tools.CopyFile("C:/Users/dell/go/src/tool_update_service/files"+fileName, newDir+"/"+fileName)
			tarFile := newDir + "/" + fileName
			sourceFile := "./files" + fileName
			tools.CopyFile(sourceFile, tarFile)
			nMd5 := fileInfo.CalcMd5FromPath(tarFile)
			zap.L().Info(nMd5)
			if len(nMd5) <= 0 {
				zap.L().Error("计算新文件md5值失败，请确保更新后，文件成功归类！", zap.Error(errors.New("calc_file_md5_fialed")))
			}
			if nMd5 == p.MD5 {
				err := os.Remove(sourceFile)
				if err != nil {
					zap.L().Error("删除源文件失败！", zap.Error(err))
				}
			} else {
				zap.L().Error("新文件与源文件不一致，请确保更新后，文件成功归类！", zap.Error(errors.New("new file does not match the source file!")))
			}
		}()
	}

	u.Url = fmt.Sprintf("%s%s%s", url_base, newDir, fileName)
	fmt.Printf("文件的url_base_ffff：%v\n", u.Url)
	//4.写入版本更新信息
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
