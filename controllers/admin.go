package controllers

type AdminController struct {
	BaseController
}

func (this *AdminController) Get() {
	// o := orm.NewOrm()
	// var maps []orm.Params
	// num, err :=o.QueryTable("ky_admin").Filter("Id", 1).Values(&maps , "id" , "email")
	// if err == nil {
	// 	fmt.Printf("Result Nums: %d\n", num)
	// 	for _, m := range maps {
	// 		logs.SetLogger("————————")
	// 		this.OutSuccess(m)
	// 	}
	// }

}
