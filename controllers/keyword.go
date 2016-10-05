package controllers

import (

	"github.com/astaxie/beego"

	"beeblog/models"

	"strconv"
	"time"
	"strings"
	"fmt"

)


type PagerA struct {
	TotalCount int64
	//总页数
	PageCount int64
	//每页显示的条数
	PageSize int64
	//当前页
	Page int64

	//前三页
	PrevPages    []int64
	//后三页
	NextPages    []int64


}


type PagerC struct {
	PagerA

	Items  []*models.As_customs



}
type PagerB struct {
	PagerA
	Items  []*models.As_keywords



}

func (this *KeywordController) ModifyCustomStatus() {
	customid:=this.Input().Get("custom.id")

	var CUstomiddd int64
	var err error
	if len(customid)>0 {
		CUstomiddd,err=strconv.ParseInt(customid,10,64)
		if err != nil {
			beego.Error(err)
		}

	}

	customStatus:=this.Input().Get("custom.customStatus")
	if customStatus=="1" {

		custom:=models.As_customs{Id:CUstomiddd,CustomStatus:0}
		err=custom.Update("CustomStatus")

		if err != nil {
			beego.Error(err)
		}



	}else if customStatus=="0"  {

		custom:=models.As_customs{Id:CUstomiddd,CustomStatus:1}
		err=custom.Update("CustomStatus")

		if err != nil {
			beego.Error(err)
		}
	}



	this.Ctx.Redirect(302,"/customlist")



}
type KeywordController struct {
	LoginController
}



func (this *KeywordController) Keywordxufei() {

	timee:=time.Now()


	result:="exception"

	iddd:=this.Input().Get("keywords.id")

	idddd,err:=strconv.ParseInt(iddd,10,64)

	if err != nil {
		beego.Error(err)
	}



	var price float64
	p:=this.Input().Get("p")
	pp:=p
	var serviceTypeId int64
	var serviceYear int64



	if(strings.Contains(p,"-")){
		dizhi1:=strings.Split(p,"-")

		p=dizhi1[0]
		var err error
		serviceTypeId,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		p=dizhi1[1]
	}
	systemconfig:=models.As_systemconfig{Id:serviceTypeId}
	err=systemconfig.Read("Id")
	if err != nil {
		beego.Error(err)
	}
	stt:=systemconfig.ConfigValue
	st,err:=strconv.ParseInt(stt,10,64)
	if err != nil {
		beego.Error(err)
	}
	if(!strings.Contains(p,"id_")){

		serviceYear,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		pric:=serviceYear*st
		pri:=strconv.FormatInt(pric,10)
		price,err=strconv.ParseFloat(pri,10)
		if err != nil {
			beego.Error(err)
		}

	}else {


		list := strings.Split(p, "_")
		Idd:=list[1]
		Id,err:= strconv.ParseInt(Idd,10,64)
		if err != nil {
			beego.Error(err)
		}else {
			systemconfigssss:=models.As_systemconfig{Id:Id}
			err:=systemconfigssss.Read("Id")
			if err != nil {
				beego.Error(err)
			}
			yearssss:=systemconfigssss.ConfigValue
			serviceYear,err= strconv.ParseInt(yearssss,10,64)
			pric:=serviceYear*st



			pri:=strconv.FormatInt(pric,10)
			price,err=strconv.ParseFloat(pri,10)
			if err != nil {
				beego.Error(err)
			}

		}




	}


	user:=GetCurrentUser()
	account:=models.As_account{User:user}
	account.Read("User")


	if err != nil {
		beego.Error(err)
	}


	if account.Money>=price {



		dizhi1:=strings.Split(pp,"-")

		p=dizhi1[0]
		var err error
		serviceTypeId,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		p=dizhi1[1]





		if (strings.Contains(p,"id_")){


			list := strings.Split(p, "_")
			Idd:=list[1]
			Idssss,err:= strconv.ParseInt(Idd,10,64)
			if err != nil {
				beego.Error(err)
			}else {
				systemconfigssss:=models.As_systemconfig{Id:Idssss}
				err:=systemconfigssss.Read("Id")
				if err != nil {
					beego.Error(err)
				}
				aaString:=systemconfigssss.ConfigTypeName

				ss := strings.Split(aaString, "赠")
				zengString:=ss[1]
				zeng,err:=strconv.ParseInt(zengString,10,64)
				if err != nil {
					beego.Error(err)
				}
				ConfigValue,err:= strconv.ParseInt(systemconfigssss.ConfigValue,10,64)
				if err != nil {
					beego.Error(err)
				}
				serviceYear=ConfigValue+zeng


			}


		}



		PassDateTime:=timee

		Year:=strconv.FormatInt(serviceYear,10)
		years,err:=strconv.Atoi(Year)
		if err != nil {
			beego.Error(err)
		}







		PassDateTime=PassDateTime.AddDate(years,0,0)

		keywordssss:=models.As_keywords{
			RegPassDatetime:PassDateTime,
			Id:idddd}




		flagg:=models.Tx_SaveKeywords(&keywordssss,user)

		if flagg {
			result="success"
		}else {
			result="failed"
		}







	}else {
		result="nomoney"
	}




	this.Ctx.WriteString(result)



}

func (this *KeywordController) Xufei() {

	id := this.Input().Get("keywords.id")


	if len(id)>0  {


		iddd,err:=strconv.ParseInt(id,10,64)
		if err != nil {
			beego.Error(err)
		}

		keywords:=models.As_keywords{Id:iddd}
		err=keywords.Read()
		if err != nil {
			beego.Error(err)
		}


		this.Data["keywords"]=keywords

	}





	currentUser:= GetCurrentUser()

	this.Data["currentUser"] = currentUser





	a:= strconv.FormatInt(currentUser.Id,10)

	account,err:=models.GetAccountByUserId(a)
	if err != nil {
		beego.Error(err)
	}
	this.Data["account"] = account


	this.Data["roleFunctions"] = GetRoleList()



	this.Data["serviceType"]=GetSystemConfig()



	this.TplName = "xufei.html"


}
func (this *KeywordController) Modifyapp() {
	result:="shibai"

	id := this.Input().Get("keywords.id")
	appUserName := this.Input().Get("keywords.appUserName")
	appPassword := this.Input().Get("keywords.appPassword")


	if len(id)>0  {


		iddd,err:=strconv.ParseInt(id,10,64)
		if err != nil {
			beego.Error(err)
		}

		keywords:=models.As_keywords{Id:iddd,OpenApp:1,AppPassword:appPassword,AppUserName:appUserName}
		err=keywords.Update("OpenApp","AppPassword","AppUserName")
		if err != nil {
			beego.Error(err)

		}else {
			result="success"
		}





	}

	this.Ctx.WriteString(result)

}

func (this *KeywordController) Openapp() {
	id := this.Input().Get("keywords.id")


	if len(id)>0  {


		iddd,err:=strconv.ParseInt(id,10,64)
		if err != nil {
			beego.Error(err)
		}

		keywords:=models.As_keywords{Id:iddd}
		err=keywords.Read()
		if err != nil {
			beego.Error(err)
		}

		systemConfig:=models.As_systemconfig{Id:keywords.ProductType}
		err=systemConfig.Read("Id")
		if err != nil {
			beego.Error(err)
		}

		this.Data["systemConfig"]=systemConfig

		this.Data["keywords"]=keywords

	}






	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()
	this.TplName = "openapp.html"





}


func (this *KeywordController) Delectkeyword() {
	result:="shibai"

	id := this.Input().Get("keywords.id")
	if len(id)>0  {


	iddd,err:=strconv.ParseInt(id,10,64)
	if err != nil {
		beego.Error(err)
	}

		keywords:=models.As_keywords{Id:iddd}
		err=keywords.Delete()
		if err != nil {
			beego.Error(err)
		}else {
			result="success"
		}
	}
	this.Ctx.WriteString(result)


}
func (this *KeywordController) Updatekeywords() {
	user:=GetCurrentUser()
	result:="shibai"
	id := this.Input().Get("keywords.id")
	checkStatus := this.Input().Get("keywords.checkStatus")
	isUse := this.Input().Get("keywords.isUse")

	if len(id)>0 && len(isUse)>0 {
		iddd,err:=strconv.ParseInt(id,10,64)
		if err != nil {
			beego.Error(err)
		}

		isUsessss,err:=strconv.ParseInt(isUse,10,64)
		if err != nil {
			beego.Error(err)
		}


		keywords:=models.As_keywords{Id:iddd,IsUse:isUsessss}



		err=keywords.Update("IsUse")
		if err != nil {
			beego.Error(err)
		}

		result="success"




	}



	if len(id)>0 && len(checkStatus)>0 {
		iddd,err:=strconv.ParseInt(id,10,64)
		if err != nil {
			beego.Error(err)
		}

		checkStatussss,err:=strconv.ParseInt(checkStatus,10,64)
		if err != nil {
			beego.Error(err)
		}

		keywords:=models.As_keywords{Id:iddd,CheckStatus:checkStatussss}
		if checkStatussss==1 {

			err=keywords.Update("CheckStatus")
			if err != nil {
				beego.Error(err)
			}

			result="success"
		}else if  checkStatussss==2  {
			flagg:= models.Tx_ChangeStatusToOk(&keywords,user)

			if flagg {
				result="success"
			}

		}else if  checkStatussss==3  {
			flagg:= models.Tx_ChangeStatusToNo(&keywords,user)

			if flagg {
				result="success"
			}

		}

		
		
		
		
		
		



	}
	
	
this.Ctx.WriteString(result)

}


func (this *KeywordController) Checkkeyword() {


	var Page1 PagerB
	var num int64
	num=3

	Page1.PageSize=10


	keywords := this.Input().Get("keywords")



	//先声明一个string 在接收Value

	if len(keywords)==0{
		keywords=""
	}
	this.Data["keywords"]=keywords









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

	count, UserList,err:=models.GetKeywordBySearch(keywords,StarNu,yedax)
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

	this.TplName = "checkkeywords.html"




}



func (this *KeywordController) Keywordmanage() {



	var Page1 PagerB
	var num int64
	num=3

	Page1.PageSize=10


	keywords := this.Input().Get("keywords")



	//先声明一个string 在接收Value

	if len(keywords)==0{
		keywords=""
	}
	this.Data["keywords"]=keywords









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

	count, UserList,err:=models.GetKeywordBySearch(keywords,StarNu,yedax)
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

	this.TplName = "keywordmanage.html"





}
func (this *KeywordController) Modifysavecustom() {

	customid:=this.Input().Get("custom.id")
var CUstomiddd int64
	var err error
	if len(customid)>0 {
		CUstomiddd,err=strconv.ParseInt(customid,10,64)
		if err != nil {
			beego.Error(err)
		}

	}


	customcustomName:=this.Input().Get("custom.customName")
	customcustomTypeName:=this.Input().Get("custom.customTypeName")

	customcustomType:=this.Input().Get("custom.customType")
	var customType int64
	var erree error
	if len(customcustomType)>0 {
		customType,erree=strconv.ParseInt(customcustomType,10,64)
		if erree != nil {
			beego.Error(erree)
		}

	}

	siteUrl:=this.Input().Get("custom.siteUrl")

	customStatus:=this.Input().Get("custom.customStatus")
	var customStatuss int64
	var erre error
	if len(customcustomType)>0 {
		customStatuss,erre=strconv.ParseInt(customStatus,10,64)
		if erre != nil {
			beego.Error(erre)
		}

	}


	bossName:=this.Input().Get("custom.bossName")


	cardTypeName:=this.Input().Get("custom.cardTypeName")

	cardType:=this.Input().Get("custom.cardType")
	var cardTypes int64
	var errr error
	if len(customcustomType)>0 {
		cardTypes,errr=strconv.ParseInt(cardType,10,64)
		if errr != nil {
			beego.Error(errr)
		}

	}


	cardNum:=this.Input().Get("custom.cardNum")

	country:=this.Input().Get("custom.country")

	province:=this.Input().Get("custom.province")

	companyFax:=this.Input().Get("custom.companyFax")

	city:=this.Input().Get("custom.city")

	companyTel:=this.Input().Get("custom.companyTel")

	area:=this.Input().Get("custom.area")

	companyAddress:=this.Input().Get("custom.companyAddress")

	memo:=this.Input().Get("custom.memo")

	As_customs:=models.As_customs{Id:CUstomiddd,CustomName:customcustomName,CustomTypeName:customcustomTypeName,CustomType:customType,SiteUrl:siteUrl,CustomStatus:customStatuss,
		BossName:bossName,RegDatetime:time.Now(),CardTypeName:cardTypeName,CardType:cardTypes,CardNum:cardNum,Country:country,
		Province:province,CompanyFax:companyFax,City:city,CompanyTel:companyTel,Area:area,CompanyAddress:companyAddress,
		Memo:memo,Agent:GetCurrentUser()}






	contactcount:=this.Input().Get("contactcount")


	var count int
	var errrr error
	if len(contactcount)>0 {

		count,errrr=strconv.Atoi(contactcount)
	}
	contactsList:=make([]*models.As_contacts,0)

	if errrr==nil&&count>-1{




		for  i:=0;i<=count;i++ {
			contactNamesssssss:=""
			contactTelssssssss:=""
			contactFaxss:=""
			contactEmailss:=""
			contactRoless:=""
			ii:=strconv.Itoa(i)
			contactNamesssssss="contactsList["+ii+"].contactName"
			contactTelssssssss="contactsList["+ii+"].contactTel"
			contactFaxss="contactsList["+ii+"].contactFax"
			contactEmailss="contactsList["+ii+"].contactEmail"
			contactRoless="contactsList["+ii+"].contactRole"
			contactName:=this.Input().Get(contactNamesssssss)
			contactTel:=this.Input().Get(contactTelssssssss)
			contactFax:=this.Input().Get(contactFaxss)
			contactEmail:=this.Input().Get(contactEmailss)
			contactRole:=this.Input().Get(contactRoless)
			contact:= models.As_contacts{ContactName:contactName,ContactTel:contactTel,ContactEmail:contactEmail,ContactRole:contactRole,ContactFax:contactFax}
			contactsList=append(contactsList,&contact)

		}




		flag:=models.Tx_ModifyCustomContact(contactsList,&As_customs)

		if flag {
			this.Ctx.Redirect(302,"/customlist")
		}else {
			this.Ctx.Redirect(302,"/keyword_addcustom")
		}


	}else {
		err:=As_customs.Insert()

		if err!=nil {
			beego.Error(err)
			this.Ctx.Redirect(302,"/keyword_addcustom")

		}else {
			this.Ctx.Redirect(302,"/customlist")
		}
	}



}

func (this *KeywordController) Modifycustom() {

	Id := this.Input().Get("custom.id")
	idd,err:=strconv.ParseInt(Id,10,64)
	if err != nil {
		beego.Error(err)
	}
	customs:=models.As_customs{Id:idd}
	customs.Read("Id")
	constact,err:=models.GetConstactByCumstom(&customs)
	if err != nil {
		beego.Error(err)
	}

	areaId:= customs.Area


	areassa,err:=models.GetAreaById(areaId)
	if err != nil {
		beego.Error(err)
	}
	this.Data["area"]=areassa


	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()


	this.Data["custom"]=customs
	this.Data["contactsList"]=constact
  aa:=len(constact)
	this.Data["contactcount"]=aa

	customTypeList,err:=models.GetSystemconfigByConfigType(5)
	if err != nil {
		beego.Error(err)
	}

	this.Data["customTypeList"]=customTypeList



	cardTypeList,err:=models.GetSystemconfigByConfigType(6)
	if err != nil {
		beego.Error(err)
	}



	this.Data["cardTypeList"]=cardTypeList

	provinceList,err:=models.GetprovinceList()
	if err != nil {
		beego.Error(err)
	}

	this.Data["provinceList"]=provinceList




	this.TplName = "modifycustom.html"

}

func (this *KeywordController) Chakancustom() {

	Id := this.Input().Get("custom.id")
	idd,err:=strconv.ParseInt(Id,10,64)
	if err != nil {
		beego.Error(err)
	}
	 customs:=models.As_customs{Id:idd}
	customs.Read("Id")
	constact,err:=models.GetConstactByCumstom(&customs)
	if err != nil {
		beego.Error(err)
	}

	 areaId:= customs.Area


	areassa,err:=models.GetAreaById(areaId)
	if err != nil {
		beego.Error(err)
	}
 	 this.Data["area"]=areassa


	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()


	this.Data["custom"]=customs
	this.Data["contactsList"]=constact


	this.TplName = "chakancustom.html"



}
func (this *KeywordController) Customlistsearch() {

	var Page1 PagerC
	var num int64
	num=3

	Page1.PageSize=10


	name := this.Input().Get("customsname")



	//先声明一个string 在接收Value

	if len(name)==0{
		name=""
	}
	this.Data["customsname"]=name









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

	count, UserList,err:=models.GetCustomByCustomName(name,StarNu,yedax)
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

	this.TplName = "customlist.html"

}
func (this *KeywordController) Account() {
	currentUser:= GetCurrentUser()

	this.Data["currentUser"] = currentUser





	a:= strconv.FormatInt(currentUser.Id,10)

	account,err:=models.GetAccountByUserId(a)
	if err != nil {
		beego.Error(err)
	}
	Money:= account.Money
	m:=strconv.FormatFloat(Money,'f',2,64)
	this.Ctx.WriteString(m)
}
func (this *KeywordController) Submitkeyword() {
	timee:=time.Now()


	result:="exception"



	var price float64
	p:=this.Input().Get("p")
	pp:=p
	var serviceTypeId int64
	var serviceYear int64



	if(strings.Contains(p,"-")){
		dizhi1:=strings.Split(p,"-")

		p=dizhi1[0]
		var err error
		serviceTypeId,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		p=dizhi1[1]
	}
	systemconfig:=models.As_systemconfig{Id:serviceTypeId}
	err:=systemconfig.Read("Id")
	if err != nil {
		beego.Error(err)
	}
	stt:=systemconfig.ConfigValue
	st,err:=strconv.ParseInt(stt,10,64)
	if err != nil {
		beego.Error(err)
	}
	if(!strings.Contains(p,"id_")){

		serviceYear,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		pric:=serviceYear*st
		pri:=strconv.FormatInt(pric,10)
		price,err=strconv.ParseFloat(pri,10)
		if err != nil {
			beego.Error(err)
		}

	}else {


		list := strings.Split(p, "_")
		Idd:=list[1]
		Id,err:= strconv.ParseInt(Idd,10,64)
		if err != nil {
			beego.Error(err)
		}else {
			systemconfigssss:=models.As_systemconfig{Id:Id}
			err:=systemconfigssss.Read("Id")
			if err != nil {
				beego.Error(err)
			}
			yearssss:=systemconfigssss.ConfigValue
			serviceYear,err= strconv.ParseInt(yearssss,10,64)
			pric:=serviceYear*st



			pri:=strconv.FormatInt(pric,10)
			price,err=strconv.ParseFloat(pri,10)
			if err != nil {
				beego.Error(err)
			}

		}




	}

	keywords:=this.Input().Get("keywords.keywords")
	customId:=this.Input().Get("keywords.customId")
	customName:=this.Input().Get("keywords.customName")
	user:=GetCurrentUser()
	account:=models.As_account{User:user}
	account.Read("User")
	sttt:=strconv.FormatInt(st,10)
	sttttt,err:=strconv.ParseFloat(sttt,10)
	if err != nil {
		beego.Error(err)
	}


	if account.Money>=price {



		dizhi1:=strings.Split(pp,"-")

		p=dizhi1[0]
		var err error
		serviceTypeId,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		p=dizhi1[1]





	if (strings.Contains(p,"id_")){


		list := strings.Split(p, "_")
		Idd:=list[1]
		Idssss,err:= strconv.ParseInt(Idd,10,64)
		if err != nil {
			beego.Error(err)
		}else {
			systemconfigssss:=models.As_systemconfig{Id:Idssss}
			err:=systemconfigssss.Read("Id")
			if err != nil {
				beego.Error(err)
			}
			aaString:=systemconfigssss.ConfigTypeName

			ss := strings.Split(aaString, "赠")
			zengString:=ss[1]
			zeng,err:=strconv.ParseInt(zengString,10,64)
			if err != nil {
				beego.Error(err)
			}
			ConfigValue,err:= strconv.ParseInt(systemconfigssss.ConfigValue,10,64)
			if err != nil {
				beego.Error(err)
			}
			serviceYear=ConfigValue+zeng


		}


	}



		PassDateTime:=timee

		Year:=strconv.FormatInt(serviceYear,10)
		years,err:=strconv.Atoi(Year)
		if err != nil {
			beego.Error(err)
		}
		id,err:=strconv.ParseInt(customId,10,64)
		if err != nil {
			beego.Error(err)
		}
		custom:=models.As_customs{Id:id}
		fmt.Println("years",years)

fmt.Println("serviceYear",serviceYear)
		PassDateTime=PassDateTime.AddDate(years,0,0)
		fmt.Println("PassDateTime",PassDateTime)
		keywordssss:=models.As_keywords{Agent:user,AgentName:user.UserName,PreRegFrozenMoney:sttttt,
		Price:price,ProductType:systemconfig.Id,ServiceYears:serviceYear,
PreRegDatetime:timee,PreRegPassDatetime:timee,RegDatetime:timee,
RegPassDatetime:PassDateTime,IsPass:0,IsUse:1,CheckStatus:0,OpenApp:0,
			Keywords:keywords,Custom:&custom,CustomName:customName}




		flagg:=models.Tx_SaveKeywords(&keywordssss,user)

		if flagg {
			result="success"
		}else {
			result="failed"
		}







	}else {
		result="nomoney"
	}




	this.Ctx.WriteString(result)

}
func (this *KeywordController) Jisuan() {
	result:="exception"
	p:=this.Input().Get("p")
var serviceTypeId int64
	var serviceYear int64



	if(strings.Contains(p,"-")){
		dizhi1:=strings.Split(p,"-")

		p=dizhi1[0]
		var err error
		serviceTypeId,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		p=dizhi1[1]
	}
	systemconfig:=models.As_systemconfig{Id:serviceTypeId}
	err:=systemconfig.Read("Id")
	if err != nil {
		beego.Error(err)
	}
	stt:=systemconfig.ConfigValue
st,err:=strconv.ParseInt(stt,10,64)
	if err != nil {
		beego.Error(err)
	}
	if(!strings.Contains(p,"id_")){

		serviceYear,err=strconv.ParseInt(p,10,64)
		if err != nil {
			beego.Error(err)
		}
		re:=serviceYear*st
		result=strconv.FormatInt(re,10)


	}else {


		list := strings.Split(p, "_")
		Idd:=list[1]
		Id,err:= strconv.ParseInt(Idd,10,64)
		if err != nil {
			beego.Error(err)
		}else {
			systemconfigssss:=models.As_systemconfig{Id:Id}
			err:=systemconfigssss.Read("Id")
			if err != nil {
				beego.Error(err)
			}
			yearssss:=systemconfigssss.ConfigValue
			serviceYear,err= strconv.ParseInt(yearssss,10,64)
			reeee:=serviceYear*st
			result=strconv.FormatInt(reeee,10)

		}




	}




	this.Ctx.WriteString(result)
}

func (this *KeywordController) Deleteconfig() {
	id:=this.Input().Get("systemConfig.id")
	idss,err:=strconv.ParseInt(id,10,64)
	if err != nil {
		beego.Error(err)
	}


sy:=models.As_systemconfig{Id:idss}
	err=sy.Delete()
	if err != nil {
		beego.Error(err)
	}else {
		this.Ctx.WriteString("11")
	}
}
func (this *KeywordController) Addconfig() {

	configType:=this.Input().Get("systemConfig.configType")
	configTypeName:=this.Input().Get("systemConfig.configTypeName")
	isStart:=this.Input().Get("systemConfig.isStart")
	configValue:=this.Input().Get("systemConfig.configValue")
	i,err:=models.IsCunZaiConfigTypeName(configTypeName)
	if err != nil {
		beego.Error(err)
	}else {
		if i>0 {
			this.Ctx.WriteString("peat")
		}else {
			configTypess,err:=strconv.ParseInt(configType,10,64)
			if err != nil {
				beego.Error(err)
			}

			isStartss,err:=strconv.ParseInt(isStart,10,64)
			if err != nil {
				beego.Error(err)
			}

			systemConfig:=models.As_systemconfig{ConfigType:configTypess,ConfigTypeName:configTypeName,IsStart:isStartss,ConfigValue:configValue}
			err=systemConfig.Insert()
			if err != nil {
				beego.Error(err)
			}else {
				this.Ctx.WriteString("noopeat")
			}


		}
	}


}
func (this *KeywordController) Youhuitype() {

	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()
	this.Data["configType"] =7


	customTypeList,err:=models.GetSystemconfigByConfigType(7)
	if err != nil {
		beego.Error(err)
	}

	this.Data["systemConfigList"]=customTypeList




	this.TplName = "systemconfig.html"
}


func (this *KeywordController) Keyword_valikey() {
	customName:=this.Input().Get("keywords.keywords")
	i,err:=models.IsCunZaikeywords(customName)
	if err != nil {
		beego.Error(err)
	}
	if i>0 {
		this.Ctx.WriteString("nosuccess")
	}else {
		this.Ctx.WriteString("success")
	}


}

func (this *KeywordController) Isexitcustomname() {
	 customName:=this.Input().Get("custom.customName")

	i,err:=models.IsCunZaicustem(customName)
	if err != nil {
		beego.Error(err)
	}
	if i>0 {
		this.Ctx.WriteString("peat")
	}else {
		this.Ctx.WriteString("nopeat")
	}

}

func (this *KeywordController) Addsavecustom() {

	customcustomName:=this.Input().Get("custom.customName")
	customcustomTypeName:=this.Input().Get("custom.customTypeName")

	customcustomType:=this.Input().Get("custom.customType")
	var customType int64
	var erree error
	if len(customcustomType)>0 {
		customType,erree=strconv.ParseInt(customcustomType,10,64)
		if erree != nil {
			beego.Error(erree)
		}

	}

	siteUrl:=this.Input().Get("custom.siteUrl")

	customStatus:=this.Input().Get("custom.customStatus")
	var customStatuss int64
	var erre error
	if len(customcustomType)>0 {
		customStatuss,erre=strconv.ParseInt(customStatus,10,64)
		if erre != nil {
			beego.Error(erre)
		}

	}


	bossName:=this.Input().Get("custom.bossName")


	cardTypeName:=this.Input().Get("custom.cardTypeName")

	cardType:=this.Input().Get("custom.cardType")
	var cardTypes int64
	var errr error
	if len(customcustomType)>0 {
		cardTypes,errr=strconv.ParseInt(cardType,10,64)
		if errr != nil {
			beego.Error(errr)
		}

	}


	cardNum:=this.Input().Get("custom.cardNum")

	country:=this.Input().Get("custom.country")

	province:=this.Input().Get("custom.province")

	companyFax:=this.Input().Get("custom.companyFax")

	city:=this.Input().Get("custom.city")

	companyTel:=this.Input().Get("custom.companyTel")

	area:=this.Input().Get("custom.area")

	companyAddress:=this.Input().Get("custom.companyAddress")

	memo:=this.Input().Get("custom.memo")

	As_customs:=models.As_customs{CustomName:customcustomName,CustomTypeName:customcustomTypeName,CustomType:customType,SiteUrl:siteUrl,CustomStatus:customStatuss,
	BossName:bossName,RegDatetime:time.Now(),CardTypeName:cardTypeName,CardType:cardTypes,CardNum:cardNum,Country:country,
	Province:province,CompanyFax:companyFax,City:city,CompanyTel:companyTel,Area:area,CompanyAddress:companyAddress,
	Memo:memo,Agent:GetCurrentUser()}






	contactcount:=this.Input().Get("contactcount")


	var count int
	var errrr error
	if len(contactcount)>0 {

		count,errrr=strconv.Atoi(contactcount)
	}
	contactsList:=make([]*models.As_contacts,0)

	if errrr==nil&&count>-1{




		for  i:=0;i<=count;i++ {
			contactNamesssssss:=""
			contactTelssssssss:=""
			contactFaxss:=""
			contactEmailss:=""
			contactRoless:=""
			ii:=strconv.Itoa(i)
			contactNamesssssss="contactsList["+ii+"].contactName"
			contactTelssssssss="contactsList["+ii+"].contactTel"
			contactFaxss="contactsList["+ii+"].contactFax"
			contactEmailss="contactsList["+ii+"].contactEmail"
			contactRoless="contactsList["+ii+"].contactRole"
			contactName:=this.Input().Get(contactNamesssssss)
			contactTel:=this.Input().Get(contactTelssssssss)
			contactFax:=this.Input().Get(contactFaxss)
			contactEmail:=this.Input().Get(contactEmailss)
			contactRole:=this.Input().Get(contactRoless)
			contact:= models.As_contacts{ContactName:contactName,ContactTel:contactTel,ContactEmail:contactEmail,ContactRole:contactRole,ContactFax:contactFax}
			contactsList=append(contactsList,&contact)

		}




		flag:=models.Tx_saveCustomContact(contactsList,&As_customs)

		if flag {
			this.Ctx.Redirect(302,"/customlist")
		}else {
			this.Ctx.Redirect(302,"/keyword_addcustom")
		}


	}else {
		err:=As_customs.Insert()

		if err!=nil {
			beego.Error(err)
			this.Ctx.Redirect(302,"/keyword_addcustom")

		}else {
			this.Ctx.Redirect(302,"/customlist")
		}
	}





}
func (this *KeywordController) Loadarea() {

	cityId := this.Input().Get("city.cityId")
	id,err:=strconv.ParseInt(cityId,10,64)
	if err != nil {
		beego.Error(err)
	}
	Hat_city:=models.Hat_city{Id:id}

	customlist,err:= models.GetAreAist(&Hat_city)


	if err != nil {
		beego.Error(err)
	}

	this.Data["json"]=&customlist
	this.ServeJSON()

}

func (this *KeywordController) Loadcity() {


	provinceId := this.Input().Get("province.provinceId")
	id,err:=strconv.ParseInt(provinceId,10,64)
	if err != nil {
		beego.Error(err)
	}
	Hat_province:=models.Hat_province{Id:id}

	customlist,err:= models.GetCITYList(&Hat_province)


	if err != nil {
		beego.Error(err)
	}

	this.Data["json"]=&customlist
	this.ServeJSON()


}
func (this *KeywordController) Addcustom() {

	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()


	customTypeList,err:=models.GetSystemconfigByConfigType(5)
	if err != nil {
		beego.Error(err)
	}

	this.Data["customTypeList"]=customTypeList



	cardTypeList,err:=models.GetSystemconfigByConfigType(6)
	if err != nil {
		beego.Error(err)
	}



	this.Data["cardTypeList"]=cardTypeList

	provinceList,err:=models.GetprovinceList()
	if err != nil {
		beego.Error(err)
	}

	this.Data["provinceList"]=provinceList

	this.TplName = "addcustom.html"
}

func (this *KeywordController) Customlist() {


	var Page1 PagerC
	var num int64
	num=3

	Page1.PageSize=10


	name := this.Input().Get("customsname")



	//先声明一个string 在接收Value

	if len(name)==0{
		name=""
	}
	this.Data["customsname"]=name









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

	count, UserList,err:=models.GetCustomByCustomName(name,StarNu,yedax)
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

	this.TplName = "customlist.html"
}

func (this *KeywordController) Searchcustom() {




	 customName := this.Input().Get("custom.customName")
	customlist,err:= models.GetCustomByName(customName)


	if err != nil {
		beego.Error(err)
	}

	this.Data["json"]=&customlist
	this.ServeJSON()


}
func (this *KeywordController) OperateAccount() {
	resule:="shibai"
 	accountuserId := this.Input().Get("account.userId")

	accountmoney := this.Input().Get("account.money")
	accountDetaildetailType := this.Input().Get("accountDetail.detailType")
	accountDetaildetailTypeName := this.Input().Get("accountDetail.detailTypeName")
	accountDetailmemo := this.Input().Get("accountDetail.memo")
	user:= GetCurrentUser()

	detaildetailType,errrr:=strconv.ParseInt(accountDetaildetailType,10,64)

	if errrr != nil {
		beego.Error(errrr)
		this.Ctx.WriteString(resule)
		return
	}


	oldAccount,errere:=models.GetAccountByUserId(accountuserId)
	if errere != nil {
		beego.Error(errere)
		this.Ctx.WriteString(resule)
		return
	}




	if errere==nil {
		money,errr:=strconv.ParseFloat(accountmoney,64)

		if errr != nil {
			beego.Error(errr)
			this.Ctx.WriteString(resule)
			return
		}


		newAccount:= models.As_account{Money:money}


		if len(accountDetailmemo)==0 {
			accountDetailmemo=""
		}

			timeee:=time.Now()
		accountDetail:=models.As_accountdetail{DetailTypeName:accountDetaildetailTypeName,DetailType:detaildetailType,AccountMoney:oldAccount.Money+newAccount.Money,Money:newAccount.Money,User:user,DetailDateTime:timeee,Memo:accountDetailmemo}

		info:=""

		info=user.UserCode+"对"+oldAccount.User.UserName+"进行"+accountDetail.DetailTypeName+"操作,流水金额:"+accountmoney
		log:=models.As_logs{User:user,UserName:user.UserCode,OperateInfo:info,OperateDatetime:timeee}

		flag:=models.Tx_operationAccount(oldAccount,&newAccount,&accountDetail,log)
		if flag {
			resule="success"
		}
		this.Ctx.WriteString(resule)

	}



}








func (this *KeywordController) Get() {




	fuwunianxian,err:=models.GetSystemconfigByConfigType(3)
	if err != nil {
		beego.Error(err)
	}
	var nianxian1 string
	for _,v:= range fuwunianxian{
		nianxian1= v.ConfigValue
	}

	nianxian,err:=strconv.Atoi(nianxian1)


	if err != nil {
		beego.Error(err)
	}








	list:=make([]*int,0)

	for i:=1;i<=nianxian ;i++  {
		b:=i
		list=append(list,&b)
	}


	this.Data["nianxian"] = list







	youhuiType,err:=models.GetSystemconfigByConfigType(7)
	if err != nil {
		beego.Error(err)
	}

	this.Data["youhuiType"] = youhuiType








	currentUser:= GetCurrentUser()

	this.Data["currentUser"] = currentUser





	a:= strconv.FormatInt(currentUser.Id,10)

	account,err:=models.GetAccountByUserId(a)
	this.Data["account"] = account


	this.Data["roleFunctions"] = GetRoleList()



		this.Data["serviceType"]=GetSystemConfig()




	this.TplName = "keyword.html"
}
