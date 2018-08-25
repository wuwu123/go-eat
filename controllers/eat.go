package controllers

import (
	"eat/models"
	"errors"
	"encoding/json"
	"strings"
)

type EatController struct {
	BaseController
}

func (this *EatController) Get() {
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

type EatPostData struct {
	Content string
}

// 用户注册
func (this *EatController) Post() {
	body := this.Ctx.Input.RequestBody
	var cinput models.EatContent
	json.Unmarshal(body, &cinput)
	contens := strings.Split(cinput.Content, "\n")
	for _, content := range contens {
		contentId, err := models.NewEatContent().Insert(content)
		if err == nil {
			eat := models.NewEat()
			UserId, _ := this.GetUserId()
			eat.UserID = UserId
			eat.Num = 1
			eat.ContentId = contentId
			eat.Status = models.STATUS_VALID
			eat.Insert()
		}
	}
	this.OutSuccess(map[string]interface{}{})
}
