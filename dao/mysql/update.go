package mysql

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"tool_update_service/models"
)

func VersionIsExist(v string) (err error) {
	sqlStr := `select count(version) from updateinfo where version = ?`

	var count int
	if err = db.Get(&count, sqlStr, v); err != nil {
		fmt.Printf("查询数据失败！,err:%v\n", err)
	}

	if count > 0 {
		return errors.New("versionIsExist")
	}

	return
}

func Md5IsExist(md5 string) (err error) {
	sqlStr := `select count(md5) from updateinfo where md5 = ?`
	var count int
	if err := db.Get(&count, sqlStr, md5); err != nil {
		fmt.Printf("查询数据失败！,err:%v\n", err)
	}
	if count > 0 {
		return errors.New("md5IsExist")
	}
	return
}

//var Info = new(models.UpdateInfo)

//查询最后一条数据，如果有,证明当前客户端版本跟服务端最新版本一致,则返回new，无需更新
func QueryLatestRecord() (info *models.UpdateInfo, err error) {
	sqlStr := ` select version,content,forced,status,url,create_time from updateinfo where id > ? order by id desc limit 1;`
	info = new(models.UpdateInfo)
	db.Get(info, sqlStr, 0)
	if len(info.Version) <= 0 {
		zap.L().Error("请确认历史发布流程", zap.Error(errors.New("数据库没有查询到数据...")))
		return nil, errors.New("notResult")
	}
	return
}

func QueryByVersion(v string) (err error) {
	/*sqlStr := `select count(version) from updateinfo where version = ?`
	var count int
	if err := db.Get(&count, sqlStr, v); err != nil {
		fmt.Printf("查询数据失败！,err:%v\n", err)
	}
	if count > 0 {
		return errors.New("版本已存在")
	}
	return*/
	//sqlStr := ` select version,content,forced,status,url,create_time from updateinfo where id > ? order by id desc limit 1;`
	//db.Get(Info, sqlStr, 0).Error()
	//fmt.Printf("version:%v,content:%v,forced:%v,status:%v,url:%v,create_time:%v\n", info.Version, info.Content, info.Forced, info.Status, info.Url, info.CreateTiem)
	return
}

//DB.NamedExec方法用来绑定SQL语句与map中的同名字段。
func InsertUpdateInfo(u *models.UpdateInfo) (err error) {
	sqlStr := `INSERT INTO updateinfo (version,content,forced,url,size,md5,status) VALUES (:version,:content,:forced,:url,:size,:md5,:status)`
	ret, err := db.NamedExec(sqlStr,
		map[string]interface{}{
			"version": u.Version,
			"content": u.Content,
			"forced":  u.Forced,
			"url":     u.Url,
			"size":    u.Size,
			"md5":     u.Md5,
			"status":  u.Status,
		})
	fmt.Printf("-------m=d=5:%v\n", u.Md5)
	if err != nil {
		fmt.Printf("更新信息插入失败。err:%v\n", err)
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed。err:%v\n", err)
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return
}

func QueryUpdateHistory() (u []*models.UpdateInfo, err error) {
	sqlStr := `select version,content,url,forced,size,md5,create_time from updateinfo order by id desc limit 20 `
	if err = db.Select(&u, sqlStr); err != nil {
		zap.L().Error("查询数据失败", zap.Error(err))
		return nil, err
	}
	return
}
