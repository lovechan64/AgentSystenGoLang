package controllers

import (

	"github.com/astaxie/beego"

	"beeblog/models"

	"strconv"
	"time"
)


type CaiwuController struct {
	LoginController
}



func (this *CaiwuController) OperateAccount() {
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
		/*accountDetail.setAccountMoney(oldAccount.getMoney().add(newAccount.getMoney()));
			accountDetail.setMoney(newAccount.getMoney());
			accountDetail.setUserId(newAccount.getUserId());
			accountDetail.setDetailDateTime(new Date());
			if(accountDetail.getMemo()==null){
				accountDetail.setMemo("");
			}

			Logs logs=new Logs();
			logs.setUserId(this.getCurrentUser().getId());
			logs.setUserName(this.getCurrentUser().getUserCode());
			logs.setOperateInfo(logs.getUserName()+"对"+newAccount.getUserName()+"进行"+accountDetail.getDetailTypeName()+"操作,流水金额:"+newAccount.getMoney());
			logs.setOperateDatetime(new Date());
			this.accountService.tx_operationAccount(oldAccount, newAccount, accountDetail, logs);
			this.getOut().print("success");
		 }*/

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





func (this *CaiwuController) Searchuser() {


	 userCode := this.Input().Get("user.userCode")
	  Userlist,err:= models.GetUserBySearch(userCode)


	if err != nil {
		beego.Error(err)
	}

	this.Data["json"]=&Userlist
	this.ServeJSON()


}





func (this *CaiwuController) Get() {

	this.Data["currentUser"] =  GetCurrentUser()

	this.Data["roleFunctions"] = GetRoleList()

	accountConfigList ,err:= models.GetAccountSystemconfig()

	if err != nil {
		beego.Error(err)
	}

	this.Data["accountConfigList"]=accountConfigList




	this.TplName = "caiwu.html"
}
