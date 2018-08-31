package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"eat/common"
	"errors"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

const SUCCESS = "success"
const ERROR = "error"
const TOPSESSION = "topsession"

type BaseController struct {
	beego.Controller
	ControllerName string
	ActionName     string
	User           LoginUser
}

type LoginUser struct {
	Id       int64
	Nickname string
	Openid   string
}

// 设置用户信息
func (this *BaseController) SetUser() {
	topsession := this.Ctx.Input.Header(TOPSESSION)
	redisClient := common.GetRedis()
	m, _ := redisClient.HGetAll(topsession).Result()
	this.User.Id, _ = strconv.ParseInt(m["id"], 10, 64)
	this.User.Nickname = m["nickname"]
}

func (this *BaseController) SetSession() string {
	h := md5.New()
	h.Write([]byte(this.User.Openid + "hhhh"))
	key := hex.EncodeToString(h.Sum(nil))
	redisClient := common.GetRedis()
	redisClient.HSet(key, "id", this.User.Id)
	redisClient.HSet(key, "nickname", this.User.Nickname)
	redisClient.HSet(key, "openid", this.User.Openid)
	redisClient.Close()
	return key
}

func (this *BaseController) GetUserId() (userId int64, error error) {
	return this.User.Id, error
}

// 预处理
func (this *BaseController) Prepare() {
	// 设置用户信息
	this.SetUser()
	noLogin := [3]string{"UserControllerPOST", "DefaultControllerGET", "UserControllerCode"}
	controller, action := this.GetControllerAndAction()
	rote := fmt.Sprintf("%s%s", controller, action)
	status, _ := common.Contain(rote, noLogin)
	if !status {
		if this.User.Id == 0 {
			this.OutError(make(map[string]interface{}), errors.New("无效的用户"))
		}
	}
}

// 成功输出
func (this *BaseController) OutSuccess(out map[string]interface{}) {
	returnData := make(map[string]interface{})
	returnData["data"] = out
	returnData["status"] = SUCCESS
	this.Data["json"] = returnData
	this.ServeJSON()
}

// 失败输出
func (this *BaseController) OutError(out map[string]interface{}, errorMessage error) {
	returnData := make(map[string]interface{})
	beego.Error(errorMessage)
	returnData["data"] = out
	returnData["status"] = ERROR
	returnData["message"] = errorMessage.Error()
	this.Data["json"] = returnData
	this.ServeJSON()
}
