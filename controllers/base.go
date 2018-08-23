package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"eat/common"
	"errors"
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
	Id       string
	Nickname string
}

// 设置用户信息
func (this *BaseController) SetUser() {
	topsession := this.Ctx.Input.Header(TOPSESSION)
	if topsession == "" {
		topsession = "740e3ff3f32a4daf93ef951ca270034c"
	}
	redisClient := common.GetRedis()
	m, _ := redisClient.HGetAll(topsession).Result()
	this.User.Id = m["id"]
	this.User.Nickname = m["nickname"]
}

// 预处理
func (this *BaseController) Prepare() {
	// 设置用户信息
	this.SetUser()
	noLogin := [2]string{"UserControllerPOST" , "DefaultControllerGET"}
	controller, action := this.GetControllerAndAction()
	rote := fmt.Sprintf("%s%s", controller, action)
	status, _ := common.Contain(rote, noLogin)
	if !status {
		if this.User.Id == "" {
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
