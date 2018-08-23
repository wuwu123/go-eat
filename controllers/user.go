package controllers

import (
	"encoding/json"
	"eat/models"
	"strconv"
	"errors"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	cMap := make(map[string]interface{})
	userId, err := strconv.ParseInt(this.User.Id, 10, 0)
	if err != nil {
		this.OutError(map[string]interface{}{}, err)
	}
	user, err := models.NewUser().Find(userId)
	if err != nil {
		this.OutError(map[string]interface{}{}, errors.New("用户获取失败"))
	}
	cMap["id"] = user.Id
	cMap["nickname"] = user.Nickname
	cMap["lat"] = user.Lat
	cMap["lon"] = user.Lon
	this.OutSuccess(cMap)
}

type PostInput struct {
	Nickname string
	Openid   string
	Sex      uint
}

// 用户注册
func (this *UserController) Post() {
	body := this.Ctx.Input.RequestBody
	var cinput PostInput
	json.Unmarshal(body, &cinput)
	user := models.NewUser()
	user.Openid = cinput.Openid
	user.Nickname = cinput.Nickname
	user.Sex = cinput.Sex
	if err := user.Insert(); err != nil {
		beego.Error(err)
		this.OutError(map[string]interface{}{} , errors.New("创建用户失败"))
	}
	this.OutSuccess(map[string]interface{}{})
}
