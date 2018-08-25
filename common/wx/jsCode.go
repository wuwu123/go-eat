package wx

import (
	"fmt"
	"github.com/astaxie/beego"
	"errors"
	"github.com/astaxie/beego/httplib"
)

type WxConfs struct {
	Appid  string
	Secret string
}

type CodeReturn struct {
	Openid string `json:"openid"`
	SessionKey string `json:"session_key"`
}

func WxConf() (conf WxConfs, err error) {
	conf.Appid = beego.AppConfig.String("wx_appid")
	if conf.Appid == "" {
		err = errors.New("appid 未设置")
	}
	conf.Secret = beego.AppConfig.String("wx_secret")
	if conf.Secret == "" {
		err = errors.New("secret 未设置")
	}
	return conf, err
}

func WxCode(code string) (codeRe CodeReturn, err error) {
	conf, err := WxConf()
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", conf.Appid, conf.Secret, code)
	httplib.Get(url).ToJSON(&codeRe)
	return codeRe, err
}
