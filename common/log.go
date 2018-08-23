package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"path/filepath"
	"os"
	"encoding/json"
)

// 日志注册
func RegisterLogger() {

	logs.SetLogFuncCall(true)
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)

	if beego.AppConfig.DefaultBool("log_is_async", true) {
		logs.Async(1e3)
	}
	log := beego.AppConfig.String("log_path")
	if len(log) == 0 {
		beego.Error("日志路径未设置")
	}
	logPath := filepath.Join(log, "log.log")

	if _, err := os.Stat(log); os.IsNotExist(err) {
		os.MkdirAll(log, 0777)
	}

	config := make(map[string]interface{}, 1)

	config["filename"] = logPath
	config["perm"] = "0755"
	config["rotate"] = true

	if maxLines := beego.AppConfig.DefaultInt("log_maxlines", 1000000); maxLines > 0 {
		config["maxLines"] = maxLines
	}
	if maxSize := beego.AppConfig.DefaultInt("log_maxsize", 1<<28); maxSize > 0 {
		config["maxsize"] = maxSize
	}
	if !beego.AppConfig.DefaultBool("log_daily", true) {
		config["daily"] = false
	}
	if maxDays := beego.AppConfig.DefaultInt("log_maxdays", 7); maxDays > 0 {
		config["maxdays"] = maxDays
	}
	if level := beego.AppConfig.DefaultString("log_level", "Trace"); level != "" {
		switch level {
		case "Emergency":
			config["level"] = beego.LevelEmergency;break
		case "Alert":
			config["level"] = beego.LevelAlert;break
		case "Critical":
			config["level"] = beego.LevelCritical;break
		case "Error":
			config["level"] = beego.LevelError; break
		case "Warning":
			config["level"] = beego.LevelWarning; break
		case "Notice":
			config["level"] = beego.LevelNotice; break
		case "Informational":
			config["level"] = beego.LevelInformational;break
		case "Debug":
			config["level"] = beego.LevelDebug;break
		}
	}
	b, err := json.Marshal(config)
	if  err != nil {
		beego.Error("初始化文件日志时出错 ->",err)
		beego.SetLogger("file", `{"filename":"`+ logPath + `"}`)
	}else{
		beego.SetLogger(logs.AdapterFile, string(b))
	}
	beego.SetLogFuncCall(true)
}
