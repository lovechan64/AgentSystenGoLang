package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"

)

type ZhuceController struct {
	beego.Controller
}
func (this *ZhuceController) Get() {
	// 判断是否为退出操作
/*	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}*/

	this.TplName = "zhuce.html"
}


func (this *ZhuceController) Post() {
	// 获取表单信息
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"


	err:= models.AddUser(uname,"asdas",pwd,"dasdsad","1","1")
	if err != nil {
		beego.Error(err)
	}else {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}


	this.Redirect("/", 302)
	return
}

