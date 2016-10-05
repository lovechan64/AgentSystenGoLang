package controllers

import (

	"github.com/astaxie/beego"

	"beeblog/models"
	"strconv"


)


type ZhuyeController struct {
	LoginController
}



func (this *ZhuyeController) Get() {


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
