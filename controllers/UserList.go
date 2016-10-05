package controllers

import (

	"github.com/astaxie/beego"

	"beeblog/models"
	"strconv"

	"fmt"

	"time"
	"crypto/md5"
	"encoding/hex"
)


type UserlistController struct {
	LoginController
}

type Pager struct {
	TotalCount int64
//总页数
 PageCount int64
//每页显示的条数
PageSize int64
//当前页
Page int64
	Items  []*models.As_user
//前三页
	PrevPages    []int64
//后三页
	NextPages    []int64


}

func (this *UserlistController)EditUser()  {

	typea := this.Input().Get("type")

	useruserName := this.Input().Get("user.userName")
	useruserPassword := this.Input().Get(" user.userPassword")
	roleId := this.Input().Get("user.roleId")
	 isStart := this.Input().Get("user.isStart")
	useruserCode := this.Input().Get("user.userCode")
	userid := this.Input().Get("user.id")





	h:=md5.New()
	h.Write([]byte(useruserPassword))
	useruserPassword=hex.EncodeToString(h.Sum(nil))




	if "modify"==typea {
		ssuserid,err:=strconv.ParseInt(userid,10,64)

		if err != nil {
			beego.Error(err)
		}
		roleidd,err:=strconv.ParseInt(roleId,10,64)

		if err != nil {
			beego.Error(err)
		}


		ssisStart,err:=strconv.ParseInt(isStart,10,64)

		if err != nil {
			beego.Error(err)
		}
		role:=models.As_role{Id:roleidd}

		user:=models.As_user{Id:ssuserid,UserCode:useruserCode,IsStart:ssisStart,Role:&role,UserPassword:useruserPassword,UserName:useruserName}
			err= user.EditUser()
		if err != nil {
			beego.Error(err)
			this.Ctx.WriteString("aaa")
		}
		this.Ctx.WriteString("success")
	}else if "add"==typea {
		roleidd,err:=strconv.ParseInt(roleId,10,64)

		if err != nil {
			beego.Error(err)
		}


		ssisStart,err:=strconv.ParseInt(isStart,10,64)

		if err != nil {
			beego.Error(err)
		}

	cunzai,err:=	models.IsCunZaiUserCode(useruserCode)
		if err != nil {
			beego.Error(err)
		}
		if cunzai>0 {
			this.Ctx.WriteString("repeat")
		}else {
			role:=models.As_role{Id:roleidd}




			timenow:=time.Now()

			user:=models.As_user{CreationTime:timenow,LastLoginTime:timenow,CreatedBy:GetCurrentUser().UserCode,LastUpdateTime:timenow,UserCode:useruserCode,IsStart:ssisStart,Role:&role,UserPassword:useruserPassword,UserName:useruserName}
			flag:=user.Insert()
			if !flag {

				this.Ctx.WriteString("sss")
			}else {
			this.Ctx.WriteString("success")
			}
		}

	}else if "delete"==typea {
		ssuserid,err:=strconv.ParseInt(userid,10,64)

		if err != nil {
			beego.Error(err)
		}
		user:=models.As_user{Id:ssuserid}
		flag:= user.Delete()

		if !flag {
			beego.Error(err)
			this.Ctx.WriteString("sss")
		}else {
		this.Ctx.WriteString("success")
		}
	}


}
func (this *UserlistController) Post() {
	var Page1 Pager
	var num int64
	num=3

	Page1.PageSize=10


	name := this.Input().Get("uname")



	//先声明一个string 在接收Value

	if len(name)==0{
		name=""
	}
	this.Data["uname"]=name



	userroleId := this.Input().Get("userroleId")

	if len(userroleId)==0{
		userroleId="1"
		iss,err:=	strconv.ParseInt(userroleId,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userroleId"]=iss
	}else {


		iss,err:=	strconv.ParseInt(userroleId,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userroleId"]=iss
	}







	isStart := this.Input().Get("userisStart")

	if len(isStart)==0{
		isStart="1"
		iss,err:=	strconv.ParseInt(isStart,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userisStart"]=iss

	}else {
		iss,err:=	strconv.ParseInt(isStart,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userisStart"]=iss
	}

	pagerpage := this.Input().Get("pager.page")
	var dangqianye int64
	if len(pagerpage)>0  {
		var err error
		dangqianye, err = strconv.ParseInt(pagerpage, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}else {
		dangqianye=1
	}




	Page1.Page=dangqianye








	list := make([]int64, 0)
	var font int64
	font=1
	if dangqianye>num {
		font=dangqianye-num
	}
	for i:= font; i<dangqianye;  i++{
		list=append(list,i)
	}
	Page1.PrevPages=list





	yedaxiao:=Page1.PageSize
	yedax:= strconv.FormatInt(yedaxiao,10)
	StarNum:=(dangqianye-1)*yedaxiao

	StarNu := strconv.FormatInt(StarNum, 10)
	count, UserList,err:=models.GetUserList(name,userroleId,isStart,StarNu,yedax)
	Page1.TotalCount=count








	if Page1.TotalCount%Page1.PageSize>0 {
		Page1.PageCount=Page1.TotalCount/Page1.PageSize+1
	}else {
		Page1.PageCount=Page1.TotalCount/Page1.PageSize
	}






	Page1.Items=UserList
	if err != nil {
		beego.Error(err)
	}







	listss := make([]int64, 0)
	end:=num
	if Page1.PageCount>num&&(dangqianye+num)<Page1.PageCount {
		end=dangqianye+end
	}else {
		end=Page1.PageCount
	}
	for i:=dangqianye+1;i<=end ;i++  {
		listss=append(listss,i)
	}
	Page1.NextPages=listss

	fmt.Println("当前页",dangqianye,"总计录",Page1.TotalCount,"总页数",Page1.PageSize,"开始页",Page1.Page,"前三页",Page1.PrevPages,"后三页",Page1.NextPages,"=======")
	this.Data["pager"] = Page1



	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()
	rolelist,err:=models.GetRoleList()

	this.Data["rolelist"] =rolelist

	this.TplName = "userlist.html"
}

func (this *UserlistController) Get() {
	var Page1 Pager
	var num int64
	num=3

	Page1.PageSize=10


	name := this.Input().Get("uname")



	//先声明一个string 在接收Value

	if len(name)==0{
		name=""
	}
	this.Data["uname"]=name



	userroleId := this.Input().Get("userroleId")

	if len(userroleId)==0{
		userroleId="1"
		iss,err:=	strconv.ParseInt(userroleId,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userroleId"]=iss
	}else {


		iss,err:=	strconv.ParseInt(userroleId,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userroleId"]=iss
	}







	isStart := this.Input().Get("userisStart")

	if len(isStart)==0{
		isStart="1"
		iss,err:=	strconv.ParseInt(isStart,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userisStart"]=iss

	}else {
		iss,err:=	strconv.ParseInt(isStart,10,64)
		if err != nil {
			beego.Error(err)
		}
		this.Data["userisStart"]=iss
	}

	pagerpage := this.Input().Get("pager.page")
		var dangqianye int64
	if len(pagerpage)>0  {
		var err error
		dangqianye, err = strconv.ParseInt(pagerpage, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}else {
		dangqianye=1
	}




	Page1.Page=dangqianye








	list := make([]int64, 0)
	var font int64
		font=1
	if dangqianye>num {
		font=dangqianye-num
	}
	for i:= font; i<dangqianye;  i++{
		list=append(list,i)
	}
	Page1.PrevPages=list





	yedaxiao:=Page1.PageSize
	   yedax:= strconv.FormatInt(yedaxiao,10)
	StarNum:=(dangqianye-1)*yedaxiao

	StarNu := strconv.FormatInt(StarNum, 10)
	count, UserList,err:=models.GetUserList(name,userroleId,isStart,StarNu,yedax)
	Page1.TotalCount=count








	if Page1.TotalCount%Page1.PageSize>0 {
		Page1.PageCount=Page1.TotalCount/Page1.PageSize+1
	}else {
		Page1.PageCount=Page1.TotalCount/Page1.PageSize
	}






	Page1.Items=UserList
	if err != nil {
		beego.Error(err)
	}






	listss := make([]int64, 0)
	end:=num
	if Page1.PageCount>num&&(dangqianye+num)<Page1.PageCount {
			end=dangqianye+end
	}else {
		end=Page1.PageCount
	}
	for i:=dangqianye+1;i<=end ;i++  {
		listss=append(listss,i)
	}
	Page1.NextPages=listss

     fmt.Println("当前页",dangqianye,"总计录",Page1.TotalCount,"总页数",Page1.PageSize,"开始页",Page1.Page,"前三页",Page1.PrevPages,"后三页",Page1.NextPages,"=======")
	this.Data["pager"] = Page1



	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()
			rolelist,err:=models.GetRoleList()

	this.Data["rolelist"] =rolelist

	this.TplName = "userlist.html"
}
