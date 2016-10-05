package main

import (
	//"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"beeblog/controllers"
	"beeblog/models"

	"github.com/astaxie/beego/context"

	"strconv"


	"strings"


)

func init() {
	// 注册数据库
	models.RegisterDB()
}




func main() {
	beego.BConfig.WebConfig.Session.SessionOn=true
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, false)

	var FilterUser= func(ctx *context.Context) {















		userId,ok:= ctx.Input.Session("uid").(int64)
                 cunzai:=false
		dizhi:= ctx.Request.RequestURI
		if(strings.Contains(dizhi,"_")){
			dizhi1:=strings.Split(dizhi,"_")

			dizhi=dizhi1[0]

		}

		if ok {

			Id := strconv.FormatInt(userId, 10)

		user,err:= models.GetUserById(Id)
		if err != nil {
			beego.Error(err)
		}











		Roles,err:=models.GetRoleList()
		if err != nil {
			beego.Error(err)
		}


		if err != nil {
			beego.Error(err)
		}









		FuncList := make([]*models.As_function, 0)

		for _,role :=range Roles {



			if user.Role.Id == role.Id {




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


			}
		}




			for _,v :=range FuncList{

				if dizhi==v.FuncUrl {
					cunzai=true
				}
			}




















		}
























		 if !ok&& dizhi!="/login"  &&!cunzai {
			ctx.Redirect(302,"/login")
		}


	}
	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser,true)
	// 注册 beego 路由
	beego.Router("/", &controllers.HomeController{})
/*	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")*/
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/zhuce", &controllers.ZhuceController{})
	beego.Router("/login_modifypwd", &controllers.LoginmodifypwdController{})
	beego.Router("/userlist", &controllers.UserlistController{})
	beego.Router("/userlist_edituser", &controllers.UserlistController{},"post:EditUser")
	beego.Router("/zhuye", &controllers.ZhuyeController{})
	// 附件处理
	/*
	os.Mkdir("attachment", os.ModePerm)
	beego.Router("/attachment/:all", &controllers.AttachController{})
	*/
	beego.Router("/keyword_loadcity", &controllers.KeywordController{},"post:Loadcity")
	beego.Router("/keyword_submitkeyword", &controllers.KeywordController{},"post:Submitkeyword")
	beego.Router("/keyword_account", &controllers.KeywordController{},"post:Account")




	beego.Router("/keyword_loadarea", &controllers.KeywordController{},"post:Loadarea")

	beego.Router("/keyword_addcustom", &controllers.KeywordController{},"get:Addcustom")
	beego.Router("/keyword_addsavecustom", &controllers.KeywordController{},"post:Addsavecustom")
	beego.Router("/keyword_isexitcustomname", &controllers.KeywordController{},"post:Isexitcustomname")
	beego.Router("/keywordmanage_keywordxufei", &controllers.KeywordController{},"post:Keywordxufei")





	beego.Router("/keywordmanage", &controllers.KeywordController{},"get:Keywordmanage")
	beego.Router("/checkkeyword", &controllers.KeywordController{},"get:Checkkeyword")
	beego.Router("/checkkeyword", &controllers.KeywordController{},"post:Checkkeyword")

	beego.Router("/keywordmanage", &controllers.KeywordController{},"post:Keywordmanage")

	beego.Router("/keyword_updatekeywords", &controllers.KeywordController{},"post:Updatekeywords")
	beego.Router("/keywordmanage_delectkeyword", &controllers.KeywordController{},"post:Delectkeyword")
	beego.Router("/keywordmanage_openapp", &controllers.KeywordController{},"get:Openapp")
	beego.Router("/keywordmanage_modifyapp", &controllers.KeywordController{},"post:Modifyapp")
	beego.Router("/keywordmanage_xufei", &controllers.KeywordController{},"get:Xufei")
















	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/caiwu", &controllers.CaiwuController{})
	beego.Router("/caiwu_opeaccount", &controllers.CaiwuController{},"post:OperateAccount")

	beego.Router("/keyword", &controllers.KeywordController{})
	beego.Router("/keyword_searchcustom", &controllers.KeywordController{},"post:Searchcustom")
	beego.Router("/customlist", &controllers.KeywordController{},"get:Customlist")
	beego.Router("/customlist", &controllers.KeywordController{},"post:Customlistsearch")
	beego.Router("/keyword_chakancustom", &controllers.KeywordController{},"get:Chakancustom")
	beego.Router("/keyword_modifycustom", &controllers.KeywordController{},"get:Modifycustom")
	beego.Router("/keyword_modifysavecustom", &controllers.KeywordController{},"post:Modifysavecustom")
	beego.Router("/keyword_modifyCustomStatus", &controllers.KeywordController{},"get:ModifyCustomStatus")







	beego.Router("/keyword_valikey", &controllers.KeywordController{},"post:Keyword_valikey")

	beego.Router("/youhuitype", &controllers.KeywordController{},"get:Youhuitype")
	beego.Router("/servicetype_addconfig", &controllers.KeywordController{},"post:Addconfig")
	beego.Router("/keyword_jisuan", &controllers.KeywordController{},"post:Jisuan")




	beego.Router("/servicetype_deleteconfig", &controllers.KeywordController{},"post:Deleteconfig")

	beego.Router("/caiwu_searchuser", &controllers.CaiwuController{},"post:Searchuser")
	beego.Router("/validateLoginUser", &controllers.ValidateLoginUserController{})
	// 启动 beego
	beego.Run()
}

