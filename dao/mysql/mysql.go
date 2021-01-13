package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"tool_update_service/models"
)

var db *sqlx.DB

func Init(mydb *models.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mydb.User, mydb.Password, mydb.Host, mydb.Port, mydb.Dbname)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	zap.L().Error("mysql初始化成功！")

	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
