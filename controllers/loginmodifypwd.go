package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"

	"strconv"

)

type LoginmodifypwdController struct {
	beego.Controller
}


func (this *LoginmodifypwdController) Post() {
	// 获取表单信息
	uname := this.Input().Get("user.userName")
	pwd := this.Input().Get("user.userPassword")


	ck, err := this.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
	}
	//先声明一个string 在接收Value
	name:=""
	name= ck.Value




	User,err:=models.GetUserByUserCode(name)
	if err != nil {
		beego.Error(err)
	}






		if uname == User.UserPassword {




			a:= strconv.FormatInt(User.Id,10)




			err=models.ModifyUserPassword(pwd,a)

			if err != nil {
				beego.Error(err)
			}


			this.Ctx.WriteString("success")
		}else {
			this.Ctx.WriteString("errorgggggpwd")
		}









	return
}
