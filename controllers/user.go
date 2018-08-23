package controllers

import (
	"encoding/json"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	cMap := make(map[string]interface{})
	cMap["wu_jie"] = "测试一下"
	this.OutSuccess(cMap)
}

type Postinput struct {
	Ceshi string
	Two string
}

// 用户注册
func (this *UserController) Post() {
	body := this.Ctx.Input.RequestBody
	var cinput Postinput
	json.Unmarshal(body, &cinput)
	cMap := make(map[string]interface{})
	cMap["name"] = cinput.Ceshi
	this.OutSuccess(cMap)
}
