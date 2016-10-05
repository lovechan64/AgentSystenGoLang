package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"



	"strconv"
	"crypto/md5"
	"encoding/hex"
	"time"
)
//写 var 变量  通过  Get的func来调用
var  Systemconfig  []*models.As_systemconfig
var archive   []*RoleFunctions
var currentUser *models.As_user
type RoleFunctions struct {
	MainFunction *models.As_function
	SubFunctions    []*models.As_function
}

type LoginController struct {
	beego.Controller
	RoleFunctions
	RoleFunctionList []*RoleFunctions
	LoginIn  *int64


//可以.出大写的类
	 /* RoleFunctions *[]models.As_role_premission*/
}


func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	// 获取表单信息
	uname := this.Input().Get("user.userCode")
	pwd := this.Input().Get("user.userPassword")
	yanzheng := this.Input().Get("yanzheng")
	h:=md5.New()
	h.Write([]byte(pwd))
	pwd=hex.EncodeToString(h.Sum(nil))




	if yanzheng=="1" {


		a,err:= models.IsCunZaiUserCode(uname)
		if err != nil {
			beego.Error(err)
		}

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




	}else {
	a,err:= models.IsCunZaiUserCode(uname)
	if err != nil {
		beego.Error(err)
	}

	if a>0 {


		user,err:= models.GetUserByUserCode(uname)
		if err != nil {
			beego.Error(err)
		}
		if pwd == user.UserPassword {


			err:=models.ModifyUserLast(user.Id,time.Now())
			if err != nil {
				beego.Error(err)
			}

			maxAge := 0

				currentUser=user
			this.Ctx.SetCookie("uname", uname, maxAge, "/")
			this.Ctx.SetCookie("pwd", pwd, maxAge, "/")




				this.SetSession("uid",int64(user.Id))

                




			/*
						RoleList := make([]*models.As_role, 0)
						FuncList := make([]*models.As_function, 0)

						RoleFunctionList := make([]*models.As_role_premission, 0)*/









			a:="1"

			roleFunctionList := make([]*RoleFunctions, 0)

			Roles,err:=models.GetRoleList()
			if err != nil {
				beego.Error(err)
			}
			MenuFunctions,err:=models.GetAllAs_functions(a)

			if err != nil {
				beego.Error(err)
			}










			for _,role :=range Roles {



				if user.Role.Id == role.Id {



					FuncList := make([]*models.As_function, 0)
					roleId := strconv.FormatInt(role.Id, 10)
					PremissionList, err := models.GetRolePremissionList("1", roleId)
					if err != nil {
						beego.Error(err)
					}
					for _, RolePremission := range PremissionList {
						FunctionId := strconv.FormatInt(RolePremission.FunctionId, 10)

						Function, err := models.GetFunctionById(FunctionId)
						if err != nil {
							beego.Error(err)
						}
						FuncList = append(FuncList, Function)

					}
					SubFunction := make([]*models.As_function, 0)
					for _, mf := range MenuFunctions {
						var roleFunctions	 RoleFunctions
						roleFunctions.MainFunction = mf

						if (FuncList != nil&&len(FuncList) > 0) {

							for _, f := range FuncList {

								if (f.ParentId == mf.Id) {
									SubFunction = append(SubFunction, f)

								}

							}

						}
						roleFunctions.SubFunctions = SubFunction
						SubFunction=nil
						roleFunctionList = append(roleFunctionList, &roleFunctions)

					}
					//给map属性赋值前  先用make 初始化 map属性
					/*this.Meun = make(map[int64][]*RoleFunctions)
					this.Meun[role.Id] = roleFunctionList*/

					archive= roleFunctionList


				}
			}





			Systemconfig,err= models.GetSystemconfigByConfigType(2)
			if err != nil {
				beego.Error(err)
			}




















			this.Ctx.Redirect(302,"/zhuye")


		}else {
			this.Ctx.WriteString("密码错误")
		}



	}else {


		this.Ctx.WriteString("账户不存在")


	}


	}




	return
}
func GetRoleList() []*RoleFunctions {
	return archive
}
func GetSystemConfig() []*models.As_systemconfig {
	return Systemconfig
}
func GetCurrentUser() *models.As_user {
	return currentUser
}


/*
func checkAccount(ctx *context.Context) bool {



	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}
*/



/*
func (output *LoginController) GetRoleFunctionList(ctx *context.Context) (RoleF []*RoleFunctions) {



	ck, err :=ctx.Request.Cookie("uname")


	if err != nil {
		return nil
	}

	uname := ck.Value
	user,err:= models.GetUserByUserCode(uname)
	if err != nil {
		beego.Error(err)
	}







	a:="1"

	roleFunctionList := make([]*RoleFunctions, 0)

	Roles,err:=models.GetRoleList()
	if err != nil {
		beego.Error(err)
	}
	MenuFunctions,err:=models.GetAllAs_functions(a)

	if err != nil {
		beego.Error(err)
	}

	for _,v :=range MenuFunctions{
		fmt.Print("======================",v.FunctionName)
	}









	for _,role :=range Roles {



		if user.RoleId == role.Id {



			FuncList := make([]*models.As_function, 0)
			roleId := strconv.FormatInt(role.Id, 10)
			PremissionList, err := models.GetRolePremissionList("1", roleId)
			if err != nil {
				beego.Error(err)
			}
			for _, RolePremission := range PremissionList {
				FunctionId := strconv.FormatInt(RolePremission.FunctionId, 10)

				Function, err := models.GetFunctionById(FunctionId)
				if err != nil {
					beego.Error(err)
				}
				FuncList = append(FuncList, Function)

			}
			SubFunction := make([]*models.As_function, 0)
			for _, mf := range MenuFunctions {
				var roleFunctions	 RoleFunctions
				roleFunctions.MainFunction = mf

				if (FuncList != nil&&len(FuncList) > 0) {

					for _, f := range FuncList {

						if (f.ParentId == mf.Id) {
							SubFunction = append(SubFunction, f)

						}

					}

				}
				roleFunctions.SubFunctions = SubFunction
				SubFunction=nil
				roleFunctionList = append(roleFunctionList, &roleFunctions)

			}
			//给map属性赋值前  先用make 初始化 map属性
			*/
/*this.Meun = make(map[int64][]*RoleFunctions)
			this.Meun[role.Id] = roleFunctionList*//*

			RoleF=roleFunctionList
			output.RoleFunctionList=roleFunctionList
			return roleFunctionList

		}
	}





	return RoleF
}

*/
