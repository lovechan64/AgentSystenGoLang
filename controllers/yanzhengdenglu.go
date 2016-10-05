package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"

	"fmt"
)

type ValidateLoginUserController struct {
	beego.Controller
}



func (this *ValidateLoginUserController) Post() {
	// 获取表单信息
	uname := this.Input().Get("user.userCode")
	pwd := this.Input().Get("user.userPassword")


	a,err:= models.IsCunZaiUserCode(uname)
	if err != nil {
		beego.Error(err)
	}
    fmt.Print("====================================",a)
	if a>0 {


		user,err:= models.GetUserByUserCode(uname)
		if err != nil {
			this.Ctx.WriteString("failed")
			beego.Error(err)

		}
		if pwd == user.UserPassword {
			this.Ctx.WriteString("success")
		}else {
			this.Ctx.WriteString("errorpwd")
		}



	}else {


		this.Ctx.WriteString("noexitusercode")


	}







	return
}

