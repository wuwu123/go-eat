package common

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"net/url"
	"os"
	"strings"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

//  注册数据库
func RegisterDataBase() {
	adapter := beego.AppConfig.String("db_adapter")
	if strings.EqualFold(adapter, "mysql") {
		beego.Info("初始化mysql.")
		host := beego.AppConfig.String("db_host")
		database := beego.AppConfig.String("db_database")
		username := beego.AppConfig.String("db_user")
		password := beego.AppConfig.String("db_password")
		timezone := beego.AppConfig.String("timezone")
		port := beego.AppConfig.String("db_port")
		location, err := time.LoadLocation(timezone)
		if err == nil {
			orm.DefaultTimeLoc = location
		} else {
			beego.Error("区时环境变量设置错误->", err)
		}
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s", username, password, host, port, database, url.QueryEscape(timezone))
		logs.Info(dataSource)
		if err := orm.RegisterDataBase("default", "mysql", dataSource); err != nil {
			beego.Error("mysql注册失败->", err)
			os.Exit(1)
		}
		// 最大空闲连接数
		orm.SetMaxIdleConns("default", 30)
		// 最大连接数
		orm.SetMaxOpenConns("default", 30)
	}
	beego.Info("数据库初始化完成.")
}
