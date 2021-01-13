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
	if err := db.Get(&count, sqlStr, v); err != nil {
		fmt.Printf("查询数据失败！,err:%v\n", err)
	}
	if count > 0 {
		return errors.New("版本已存在")
	}
	return
}

var Info = new(models.UpdateInfo)

//查询最后一条数据，如果有,证明当前客户端版本跟服务端最新版本一致,则返回new，无需更新
func QueryLatestRecord() (err error) {
	sqlStr := ` select version,content,forced,status,url,create_time from updateinfo where id > ? order by id desc limit 1;`
	db.Get(Info, sqlStr, 0).Error()
	if len(Info.Version) <= 0 {
		zap.L().Error("请确认发布流程", zap.Error(errors.New("数据库没有查询到数据...")))
		return errors.New("没有查询到数据")
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
	sqlStr := ` select version,content,forced,status,url,create_time from updateinfo where id > ? order by id desc limit 1;`
	db.Get(Info, sqlStr, 0).Error()
	//fmt.Printf("version:%v,content:%v,forced:%v,status:%v,url:%v,create_time:%v\n", info.Version, info.Content, info.Forced, info.Status, info.Url, info.CreateTiem)
	return
}

//DB.NamedExec方法用来绑定SQL语句与map中的同名字段。
func InsertUpdateInfo(u *models.UpdateInfo) (err error) {
	sqlStr := `INSERT INTO updateinfo (version,content,forced,url,status) VALUES (:version,:content,forced,url,status)`
	ret, err := db.NamedExec(sqlStr,
		map[string]interface{}{
			"version": u.Version,
			"content": u.Content,
			"forced":  u.Forced,
			"url":     u.Url,
			"status":  u.Status,
		})
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
