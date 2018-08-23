package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.OutSuccess(map[string]interface{}{})
}
