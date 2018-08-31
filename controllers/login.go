package controllers

import (
	"encoding/json"
	"eat/common/wx"
	"eat/models"
	"errors"
)

type LoginController struct {
	BaseController
}
type CodePost struct {
	Code string
}

// 用户菜单名称
func (this *LoginController) Code() {
	cMap := make(map[string]interface{})
	body := this.Ctx.Input.RequestBody
	var cinput CodePost
	json.Unmarshal(body, &cinput)
	ret, err := wx.WxCode(cinput.Code)
	if ret.Openid == "" {
		cMap["code"] = cinput.Code
		this.OutError(cMap, errors.New("微信调取失败"))
	}
	this.User.Openid = ret.Openid
	userModel := models.NewUser()
	userModel, _ = userModel.FindOpenid(ret.Openid)
	var UserId int64
	if userModel.Id == 0 {
		userModel.Openid = ret.Openid
		UserId, err = userModel.Insert()
		if err != nil {
			this.OutError(map[string]interface{}{}, err)
		}
	} else {
		UserId = userModel.Id
	}
	this.User.Id = UserId
	cMap["topsession"] = this.SetSession()
	this.OutSuccess(cMap)
}
