package controllers

import (
	"eat/models"
	"errors"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	cMap := make(map[string]interface{})
	user, err := models.NewUser().Find(this.User.Id)
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
// func (this *UserController) Post() {
// 	body := this.Ctx.Input.RequestBody
// 	var cinput PostInput
// 	json.Unmarshal(body, &cinput)
// 	user := models.NewUser()
// 	user.Openid = cinput.Openid
// 	user.Nickname = cinput.Nickname
// 	user.Sex = cinput.Sex
// 	if err := user.Insert(); err != nil {
// 		beego.Error(err)
// 		this.OutError(map[string]interface{}{}, errors.New("创建用户失败"))
// 	}
// 	this.OutSuccess(map[string]interface{}{})
// }

// 用户菜单名称
func (this *UserController) Eat() {
	cMap := make(map[string]interface{})
	offset, err := this.GetInt64("offset", 0)
	limit, err := this.GetInt64("limit", 10)
	eats, count, error := models.NewEat().FindByUser(this.User.Id, offset, limit)
	if error != nil {
		this.OutError(map[string]interface{}{}, err)
	}
	cMap["count"] = count
	cMap["list"] = eats
	this.OutSuccess(cMap)
}
