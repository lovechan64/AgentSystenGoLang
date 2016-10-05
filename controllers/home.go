package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
	"strconv"
)

type HomeController struct {
	LoginController
}

func (this *HomeController) Get() {

	currentUser:= GetCurrentUser()

	this.Data["currentUser"] = currentUser





	a:= strconv.FormatInt(currentUser.Id,10)

	account,err:=models.GetAccountByUserId(a)
	this.Data["account"] = account


	//不同的方法之间用  双通道来传递值

	/*this.RoleFunctionList= make([]*RoleFunctions, 0)*/


	if err != nil {
		beego.Error(err)
	}





	this.Data["roleFunctions"] = GetRoleList()





	this.TplName = "zhuye.html"
}
