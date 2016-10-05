package models

import (


	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"strconv"

	"github.com/astaxie/beego"


)


//首字母大写的表才能被外部引用
type As_account struct {
	Id              int64
	User  *As_user `orm:"rel(fk)"`
	Money float64
	MoneyBak float64
}

type As_accountdetail struct {
	Id              int64

	User  *As_user `orm:"rel(fk)"`
	DetailType       int64
	DetailTypeName string
	Money float64
	AccountMoney float64
	Memo string
	DetailDateTime time.Time
}
type As_admin struct {
	Id       int64
	UserName string
}

type As_contacts struct {
	Id       int64

	Customs  *As_customs `orm:"rel(fk)"`
	ContactName string
	ContactTel string
	ContactFax string
	ContactEmail string
	ContactRole string
}
type As_customs struct {
	Id       int64

	Agent  *As_user `orm:"rel(fk)"`
	AgentName string
	CustomName string
	CustomType int64
	CustomTypeName string
	SiteUrl string
	CustomStatus int64
	BossName string
	CardType int64
	CardTypeName string
	CardNum string
	CompanyTel string
	CompanyFax string
	RegDatetime time.Time
	Country string
	Province string
	City string
	Area string
	CompanyAddress string
	Memo string
	AgentCode string
}
type As_function struct {
	Id              int64
	FunctionCode string
	FunctionName string
	CreationTime time.Time
	CreatedBy string
	LastUpdateTime time.Time
	FuncUrl string
	IsStart int64
	ParentId int64
}

type As_keywords struct {
	Id              int64
	Keywords string
	Agent  *As_user `orm:"rel(fk)"`
	AgentName string

	Custom  *As_customs `orm:"rel(fk)"`
	CustomName string
	PreRegFrozenMoney float64
	Price float64
	ProductType int64
	ServiceYears int64
	OpenApp int64
	AppUserName string
	AppPassword  string
	LoginUrl string
	IosDownloadUrl string
	AndroidDownloadUrl string
	CodeIosUrl string
	CodeAndroidUrl string
	PreRegDatetime time.Time
	PreRegPassDatetime time.Time
	RegDatetime time.Time
	RegPassDatetime time.Time
	IsPass int64
	CheckStatus int64
	IsUse int64
}

type As_logs struct {
	Id              int64
	User  *As_user `orm:"rel(fk)"`
	UserName string
	OperateInfo string
	OperateDatetime time.Time
}



type As_role struct {
	Id              int64
	RoleName string
	CreationTime time.Time
	CreatedBy string
	LastUpdateTime time.Time
	IsStart int64
}


type As_role_premission struct {
	Id              int64
	Role *As_role `orm:"rel(fk)"`
	FunctionId int64
	CreationTime time.Time
	CreatedBy string
	LastUpdateTime time.Time
	IsStart int64
}

type As_siteconfig struct {
	Id              int64
	SiteName string
}


type As_systemconfig struct {
	Id              int64
	ConfigType int64
	ConfigTypeName string
	ConfigTypeValue int64
	ConfigValue string
	IsStart int64
}


type As_user struct {
	Id              int64
	UserCode string
	UserName string
	UserPassword string
	CreationTime time.Time
	LastLoginTime time.Time
	CreatedBy string
	LastUpdateTime time.Time
	IsStart int64
	Role *As_role `orm:"rel(fk)"`
}

type Hat_area struct {

	Are string
	Id              int64
	Area string
	City *Hat_city `orm:"rel(fk)"`
}


type Hat_city struct {
	Cit string
	Id              int64

	City string
	Province *Hat_province `orm:"rel(fk)"`
}


type Hat_province struct {

	Provi string
	Id              int64
	Province string
}
func User() orm.QuerySeter {
	return orm.NewOrm().QueryTable("as_user").OrderBy("-Id")
}




//前面的括号里面的参数表示 这个方法附属在那个类上面
func (m *As_user) EditUser()   error {

	err:=m.Update("UserCode","UserName","Role","UserPassword","IsStart")

	return err


}
func GetAccountSystemconfig()  ([]*As_systemconfig ,error ) {
	list:=make([]*As_systemconfig,0)
	o:=orm.NewOrm()
	_,err:=o.QueryTable("as_systemconfig").Filter("ConfigType",1).All(&list)




	return list,err


}
func GetprovinceList()  ([]*Hat_province ,error ) {
	list:=make([]*Hat_province,0)
	o:=orm.NewOrm()
	_,err:=o.QueryTable("hat_province").All(&list)




	return list,err


}
func GetCITYList(province *Hat_province)  ([]*Hat_city ,error ) {
	list:=make([]*Hat_city,0)
	o:=orm.NewOrm()
	_,err:=o.QueryTable("hat_city").Filter("Province",province).All(&list)




	return list,err


}

func GetAreAist(city *Hat_city)  ([]*Hat_area ,error ) {
	list:=make([]*Hat_area,0)
	o:=orm.NewOrm()
	_,err:=o.QueryTable("hat_area").Filter("City",city).All(&list)




	return list,err


}



func Tx_operationAccount(oldAccount,newAccount *As_account,accountDetail *As_accountdetail, logs As_logs) bool {
	flag:=false
	o:=orm.NewOrm()
	o.Begin()
	money:=oldAccount.Money+newAccount.Money
	oldAccount.Money=money
	oldAccount.MoneyBak=money

	 errre:=oldAccount.Update("Money","MoneyBak")


	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	}else {



		err:=accountDetail.Insert()



		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		}else {
			 eeerr:=logs.Insert()


			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			}else {
				errr:=o.Commit()
				if errr != nil {
					beego.Error(errr)
				}else {
					flag=true
				}

			}
		}


	}







return flag

}

func (m *As_customs) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *As_keywords) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *As_user) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_systemconfig) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_account) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_customs) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_keywords) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_user) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_account) Update(fields ...string) error {

	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *As_user) Insert() bool {
flagg:=false
	o := orm.NewOrm()
	var nummm  interface{}
	var list orm.ParamsList
	//自定义sql语句返回string类型的 list  查询id的最大值
	num,err:= o.Raw("select max(id) from as_user").ValuesFlat(&list)
	if  err!=nil{
		beego.Error(err)
	}


	if num >0&&err==nil{
		nummm=list[0]



	}
	var name  string
	//类型转换
	if i,ok:=nummm.(string);ok {
		name=i
	}

	num,err =strconv.ParseInt(name,10,64)
	if err != nil {
		beego.Error(err)
	}
	num=num+1

m.Id=num
	o.Begin()


  _, err = o.Insert(m);

if err != nil {
beego.Error(err)
		o.Rollback()


	}else {




			account := As_account{User:m, Money:0, MoneyBak:0}

			err := account.Insert()

			if err != nil {
				beego.Error(err)
				o.Rollback()


			} else {

				errr := o.Commit()
				if errr != nil {
					o.Rollback()
					beego.Error(errr)

				} else {
					flagg=true
				}

			}


	}

	return flagg

}
func (m *As_customs) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *As_accountdetail) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *As_logs) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}



func GetFunctionById (id string)(*As_function,error)  {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return  nil,err
	}
	o:=orm.NewOrm()
	function:=new(As_function)
	qs:=o.QueryTable("As_function")
	err=qs.Filter("id",cid).One(function)
	if err != nil {
		return nil,err
	}
	return function,nil
}
func GetRoleList()(roles []*As_role,err error){

	o:=orm.NewOrm()


	roles = make([]*As_role, 0)
	qs:=o.QueryTable("As_role")
	//   查询list的时候  需要传 数组的 内置地址作为参数
	_,err=qs.Filter("is_start",1).All(&roles)

	return roles,err

}

func GetRolePremissionList(isStart,roleId string)(role_premissions []*As_role_premission,err error){
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return nil,err
	}
	croleId, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return nil,err
	}
	o:=orm.NewOrm()


	role_premissions = make([]*As_role_premission, 0)
	qs:=o.QueryTable("As_role_premission")
	//   查询list的时候  需要传 数组的 内置地址作为参数
	_,err=qs.Filter("is_start",cisStart).Filter("role_id",croleId).All(&role_premissions)

	return role_premissions,err
}


func RegisterDB() {
	// 检查数据库文件
	orm.RegisterDriver("mysql", orm.DRMySQL)




	orm.RegisterModel(new(As_account), new(As_accountdetail),new(As_admin),new(As_contacts),new(As_customs),new(As_function),new(As_keywords),new(As_logs),new(As_role),new(As_role_premission),new(As_siteconfig),new(As_user),new(Hat_area),new(Hat_city),new(Hat_province),new(As_systemconfig))

	orm.RegisterDataBase("default", "mysql", "root:bdqn@/agentsystemdb?charset=utf8")

}




func GetAllAs_functions(isStart string) (functions []*As_function,err error){
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return nil,err
	}
	o:=orm.NewOrm()


	functions = make([]*As_function, 0)
	qs:=o.QueryTable("As_function")
	//   查询list的时候  需要传 数组的 内置地址作为参数
	_,err=qs.Filter("parent_id",0).Filter("is_start",cisStart).All(&functions)

	return functions,err
}



func AddUser(userCode,userName,userPassword,createdBy ,isStart,roleId string) error  {
	 cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return err
	}

	croleId, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return err
	}

	var role As_role
	role.Id=croleId


	o:=orm.NewOrm()
	u:=&As_user{UserCode:userCode,UserPassword:userPassword,UserName:userName,CreatedBy:createdBy,IsStart:cisStart,Role:&role,CreationTime:time.Now(),LastLoginTime:time.Now(),LastUpdateTime:time.Now()}
	_,err=o.Insert(u)
	if err != nil {
		return err

	}
	return err

}

func ModifyUserPassword(userPassword,id string) error {

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	u:=&As_user{Id:cid}
// read 从数据库看看有没有这条数据
	if o.Read(u) == nil {
		u.UserPassword=userPassword
		_,err:=o.Update(u)
		if err != nil {
			return err

		}
	}

	return nil





}

func ModifyUserLast(id int64, last  time.Time) error {


	o := orm.NewOrm()
	u:=&As_user{Id:id}
	// read 从数据库看看有没有这条数据
	if o.Read(u) == nil {
		u.LastLoginTime=last
		_,err:=o.Update(u)
		if err != nil {
			return err

		}
	}

	return nil





}

func GetUserRolename()([]*As_user,error){



	o:=orm.NewOrm()
	user:=make([]*As_user,0)
	qs:=o.QueryTable("As_user")
	_,err:=qs.Filter("Role",1).All(&user)
	if err != nil {
	return nil,err
	}
	return user,nil


}

func ModifyUser(userCode,userName,userPassword ,isStart,roleId,id string) error {
	cisStart, err := strconv.ParseInt(isStart, 10, 64)
	if err != nil {
		return err
	}

	 croleId, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		return err
	}
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	u:=&As_user{Id:cid}

	if o.Read(u) == nil {
		u.UserPassword=userPassword
		u.UserName=userName
		u.UserCode=userCode
		u.IsStart=cisStart
		 u.Role.Id=croleId
		_,err:=o.Update(u)
		if err != nil {
			return err

		}
	}

	return nil





}
func GetAccountByUserId(userId string)(*As_account,error)  {
	cid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil,err
	}
	o:=orm.NewOrm()
	account:=new(As_account)
	qs:=o.QueryTable("As_account")
	err=qs.Filter("user_id",cid).RelatedSel().One(account)
	if err != nil {
		return nil,err
	}
	return account,nil

}




func (m *As_user) Delete() bool {
  flagg:=false


	o := orm.NewOrm()
	o.Begin()
//需要new 对象来时候
	account:=new(As_account)

	 account.User=m
	errrere:=account.Read("User")
	if errrere != nil {
beego.Error(errrere)
		o.Rollback()


	}else {
		_, err := o.Delete(account)

		if err != nil {
			beego.Error(err)
			o.Rollback()


		} else {

			_, errr := o.Delete(m);


			if errr != nil {
				beego.Error(errr)
				o.Rollback()


			} else {

				errree := o.Commit()
				if errree != nil {

					beego.Error(errree)
				} else {
					flagg=true
				}

			}

		}
	}






	return flagg
}


func (m *As_keywords)Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *As_account)Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *As_systemconfig)Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *As_account) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *As_systemconfig) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}




func GetUserById(id string)(*As_user,error)  {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return  nil,err
	}
	o:=orm.NewOrm()
	user:=new(As_user)
	qs:=o.QueryTable("As_user")
	err=qs.Filter("id",cid).One(user)
	if err != nil {
		return nil,err
	}
	return user,nil

}


func GetAllUser() ([]*As_user, error) {
	o := orm.NewOrm()

	cates := make([]*As_user, 0)

	qs := o.QueryTable("As_user")
	_, err := qs.All(&cates)
	return cates, err
}


func IsCunZaiUserCode (userCode string) (int64,error) {
	o := orm.NewOrm()


	qs := o.QueryTable("As_user")
	a, err := qs.Filter("user_code",userCode).Count()
	return a, err
}
func GetUserByUserCode(userCode string) (*As_user,error)  {
	o:=orm.NewOrm()
	user:=new(As_user)
	qs:=o.QueryTable("As_user")
	err:=qs.Filter("user_code",userCode).One(user)
	if err != nil {
		return nil,err
	}
	return user,nil
}
func GetUserBySearch(usercode string) ([]*As_user, error) {
	o := orm.NewOrm()

	cates := make([]*As_user, 0)

	qs := o.QueryTable("As_user")
	_,err:=qs.Filter("user_code__contains",usercode).All(&cates)
	return cates,err
}






func GetUserList(userName ,roleId,isStart,starNum,pageSize string )  (int64,[]*As_user, error) {

	o := orm.NewOrm()

	cates := make([]*As_user, 0)

	qs := o.QueryTable("As_user")
	qs=qs.Filter("user_name__contains",userName)

	if roleId !=""{

		croleId, err := strconv.ParseInt(roleId, 10, 64)
		if err != nil {
			return 0,nil, err
		}

		qs =qs.Filter("role_id",croleId)
	}

	if isStart !="" {

		cisStart, err := strconv.ParseInt(isStart, 10, 64)
		if err != nil {
			return 0,nil, err
		}

		qs =qs.Filter("is_start",cisStart)
	}

	bb, err := qs.Count()

	if starNum !="" && pageSize!="" {

		cstarNum, err := strconv.ParseInt(starNum, 10, 64)
		if err != nil {
			return 0,nil, err
		}
		cpageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return 0,nil, err
		}
		qs =qs.Limit(cpageSize,cstarNum)
	}




//RelatedSel  自动关联查询

	_, err = qs.RelatedSel().All(&cates)

	return bb,cates, err
}


func GetSystemconfigByConfigType(ConfigType int64)([]*As_systemconfig,error)  {

	o:=orm.NewOrm()
	user:=make([]*As_systemconfig,0)
	qs:=o.QueryTable("as_systemconfig")
	_,err:=qs.Filter("ConfigType",ConfigType).All(&user)
	if err != nil {
		return nil,err
	}
	return user,nil

}

func GetCustomByName(customName string)([]*As_customs,error)  {

	o:=orm.NewOrm()
	user:=make([]*As_customs,0)
	qs:=o.QueryTable("as_customs")
	_,err:=qs.Filter("CustomName__contains",customName).All(&user)
	if err != nil {
		return nil,err
	}
	return user,nil

}



func  Tx_saveCustomContact(As_accountList []*As_contacts,customs *As_customs ) bool {




	flagg:=false
	o := orm.NewOrm()
	var nummm  interface{}
	var list orm.ParamsList
	//自定义sql语句返回list 查询id的最大值
	num,err:= o.Raw("select max(id) from as_customs").ValuesFlat(&list)
	if  err!=nil{
		beego.Error(err)
	}


		nummm=list[0]



	var name  string
	//类型转换
	if i,ok:=nummm.(string);ok {
		name=i
	}

	num,err =strconv.ParseInt(name,10,64)
	if err != nil {
		beego.Error(err)
	}
	num=num+1



	o.Begin()

	customs.Id=num

	_, err = o.Insert(customs);

	if err != nil {
		beego.Error(err)
		o.Rollback()


	}else {
		for i:=0;i<len(As_accountList) ;i++  {
			As_accountList[i].Customs=customs





		_,err := o.Insert(As_accountList[i]);

		if err != nil {
			beego.Error(err)
			o.Rollback()


		} else {

			errr := o.Commit()
			if errr != nil {
				o.Rollback()
				beego.Error(errr)

			} else {
				flagg=true
			}

		}
		}

	}

	return flagg

}



func IsCunZaicustem(customName string) (int64,error) {
	o := orm.NewOrm()


	qs := o.QueryTable("as_customs")
	a, err := qs.Filter("CustomName",customName).Count()
	return a, err
}

func IsCunZaikeywords(customName string) (int64,error) {
	o := orm.NewOrm()

	qs := o.QueryTable("As_keywords")
	a, err := qs.Filter("Keywords",customName).Count()
	return a, err
}
func IsCunZaiConfigTypeName(customName string) (int64,error) {
	o := orm.NewOrm()

	qs := o.QueryTable("As_systemconfig")
	a, err := qs.Filter("ConfigTypeName",customName).Count()
	return a, err
}


func Tx_SaveKeywords(keywords *As_keywords,user *As_user) bool {
	flag:=false
	o:=orm.NewOrm()
	o.Begin()
	_,errre:=o.Insert(keywords)

timeee:=time.Now()

	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	}else {

		account:=As_account{User:user}

		account.Read("User")
		 oney:=account.Money-keywords.Price
		account.Money=oney

account.MoneyBak=account.Money
		err:=account.Update("Money","MoneyBak")









		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		}else {
			ServiceYears:=strconv.FormatInt(keywords.ServiceYears,10)
			Price:=strconv.FormatFloat(keywords.Price,'f',2,64)

			memo:=user.UserName+"对"+keywords.CustomName+"进行关键词申请操作,扣除预付款资金："+ServiceYears+"年"+Price+"元"
			accountdetail:=As_accountdetail{DetailType:9999,User:user,DetailTypeName:"预注册冻结资金",
			DetailDateTime:timeee,Money:keywords.Price,AccountMoney:account.Money, Memo:memo }

			eeerr:=accountdetail.Insert()






			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			}else {



				log:=As_logs{User:user,UserName:user.UserCode,
				OperateDatetime:timeee,OperateInfo:memo}
				errrer:=log.Insert()
				if errrer != nil {
					beego.Error(errrer)
					o.Rollback()
					return flag

				}else {

					errr := o.Commit()
					if errr != nil {
						beego.Error(errr)
					} else {
						flag = true
					}
				}
			}
		}


	}







	return flag

}

func GetCustomByCustomName(CustomName ,starNum ,pageSize string )  (int64,[]*As_customs, error) {

	o := orm.NewOrm()

	cates := make([]*As_customs, 0)

	qs := o.QueryTable("As_customs")
	qs=qs.Filter("custom_name__contains",CustomName)


	bb, err := qs.Count()

	if starNum !="" && pageSize!="" {

		cstarNum, err := strconv.ParseInt(starNum, 10, 64)
		if err != nil {
			return 0,nil, err
		}
		cpageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return 0,nil, err
		}
		qs =qs.Limit(cpageSize,cstarNum)
	}




	//RelatedSel  自动关联查询

	_, err = qs.RelatedSel().All(&cates)

	return bb,cates, err
}

func GetConstactByCumstom(customs *As_customs)([]*As_contacts,error){



	o:=orm.NewOrm()
	user:=make([]*As_contacts,0)
	qs:=o.QueryTable("As_contacts")
	_,err:=qs.Filter("Customs",customs).All(&user)
	if err != nil {
		return nil,err
	}
	return user,nil


}



func GetAreaById(areaId string )  (*Hat_area, error) {
	areaIDd,err:=strconv.ParseInt(areaId,10,64)
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()

	cates := new(Hat_area)

	qs := o.QueryTable("hat_area")
	err=qs.Filter("Id",areaIDd).RelatedSel().One(cates)

	return cates, err
}


func  Tx_ModifyCustomContact(As_accountList []*As_contacts,customs *As_customs ) bool {




	flagg:=false
	o := orm.NewOrm()




	o.Begin()



	_, err := o.Update(customs);

	if err != nil {
		beego.Error(err)
		o.Rollback()


	}else {

		oldAccountList:=make([]*As_contacts,0)

		_,err:=o.QueryTable("As_contacts").Filter("Customs",customs).All(&oldAccountList)
		if err != nil {
			beego.Error(err)
			o.Rollback()


		}else {

			for _,v:=range oldAccountList  {
				_,err:=o.Delete(v)

				if err != nil {
					beego.Error(err)
					o.Rollback()


				}

			}














		for i:=0;i<len(As_accountList) ;i++  {
			As_accountList[i].Customs=customs





			_,err := o.Insert(As_accountList[i]);

			if err != nil {
				beego.Error(err)
				o.Rollback()


			} else {

				errr := o.Commit()
				if errr != nil {
					o.Rollback()
					beego.Error(errr)

				} else {
					flagg = true
				}
			}
			}
		}

	}

	return flagg

}

func GetKeywordBySearch(keyword ,starNum ,pageSize string )  (int64,[]*As_keywords, error) {

	o := orm.NewOrm()

	cates := make([]*As_keywords, 0)

	qs := o.QueryTable("As_keywords")
	qs=qs.Filter("keywords__contains",keyword)


	bb, err := qs.Count()

	if starNum !="" && pageSize!="" {

		cstarNum, err := strconv.ParseInt(starNum, 10, 64)
		if err != nil {
			return 0,nil, err
		}
		cpageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			return 0,nil, err
		}
		qs =qs.Limit(cpageSize,cstarNum)
	}




	//RelatedSel  自动关联查询

	_, err = qs.RelatedSel().All(&cates)

	return bb,cates, err
}

func Tx_ChangeStatusToOk(keywords *As_keywords,user *As_user) bool {
	flag:=false
  timeeee:=time.Now()


	CheckStatusssssss:=keywords.CheckStatus
	if CheckStatusssssss!=2 {
		return flag
	}




	keywords.Read("Id")

	keywords.CheckStatus=CheckStatusssssss



	o:=orm.NewOrm()
	o.Begin()

	_,errre:=o.Update(keywords,"CheckStatus")




	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	}else {

		account:= As_account{User:keywords.Agent}



		err:=account.Read("User")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		}else {


		account.Money=account.Money+keywords.PreRegFrozenMoney
		account.MoneyBak=account.Money
		_,err:=o.Update(&account,"Money","MoneyBak")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		}else {

			PreRegFrozenMoney := strconv.FormatFloat(keywords.PreRegFrozenMoney, 'f', 2, 64)

			memooo := user.UserName + "对" + keywords.CustomName + "进行关键词审核操作,返回冻结资金：" + PreRegFrozenMoney

			accountDetail := As_accountdetail{User:user, DetailType:9998,
				DetailTypeName:"返回预注册冻结资金", Money:keywords.PreRegFrozenMoney,
				AccountMoney:account.Money,
				Memo:memooo, DetailDateTime:timeeee,
			}
			_, eeerr := o.Insert(&accountDetail)

			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			} else {

				account.Money = account.Money - keywords.Price
				account.MoneyBak = account.Money

				_, err := o.Update(&account,"Money","MoneyBak")

				if err != nil {
					beego.Error(err)
					o.Rollback()
					return flag

				} else {
					memo2 := user.UserName + "对" + keywords.CustomName + "进行关键词审核通过操作自动正式扣款操作,扣除正式注册资金：" + PreRegFrozenMoney

					accountDetail2 := As_accountdetail{User:user,
						DetailType:9997, DetailTypeName:"扣除申请关键词" + keywords.Keywords + "的所有资金",
						Money:0 - account.Money,
						AccountMoney:account.Money,
						Memo:memo2,
						DetailDateTime:timeeee, }

					_, eeerr = o.Insert(&accountDetail2)
					if eeerr != nil {
						beego.Error(eeerr)
						o.Rollback()
						return flag

					} else {

						log := As_logs{User:user, UserName:user.UserCode,
							OperateDatetime:timeeee,
							OperateInfo:memooo,
						}

						_, eeerr = o.Insert(&log)
						if eeerr != nil {
							beego.Error(eeerr)
							o.Rollback()
							return flag

						} else {
							logs := As_logs{User:user, UserName:user.UserCode,
								OperateDatetime:timeeee,
								OperateInfo:memo2}


							_, eeerr = o.Insert(&logs)
							if eeerr != nil {
								beego.Error(eeerr)
								o.Rollback()
								return flag

							} else {

								errr := o.Commit()
								if errr != nil {
									beego.Error(errr)
								} else {
									flag = true
								}
							}
						}
					}
				}
				}
			}
		}

	}







	return flag

}

func Tx_ChangeStatusToNo(keywords *As_keywords,user *As_user) bool {
	flag:=false
	timeeee:=time.Now()
	CheckStatusssssss:=keywords.CheckStatus
	if CheckStatusssssss!=3 {
		return flag
	}




	keywords.Read("Id")

keywords.CheckStatus=CheckStatusssssss

	o:=orm.NewOrm()
	o.Begin()

	_,errre:=o.Update(keywords,"CheckStatus")




	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	}else {

		account:= As_account{User:keywords.Agent}

		account.Read("User")
		account.Money=account.Money+keywords.PreRegFrozenMoney
		account.MoneyBak=account.Money
		_, err := o.Update(&account,"Money","MoneyBak")

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		}else {

			PreRegFrozenMoney := strconv.FormatFloat(keywords.PreRegFrozenMoney, 'f', 2, 64)

			memooo := user.UserName + "对" + keywords.CustomName + "进行关键词审核操作,返回冻结资金：" + PreRegFrozenMoney

			accountDetail := As_accountdetail{User:user, DetailType:9998,
				DetailTypeName:"返回预注册冻结资金", Money:keywords.PreRegFrozenMoney,
				AccountMoney:account.Money,
				Memo:memooo, DetailDateTime:timeeee,
			}
			_, eeerr := o.Insert(&accountDetail)

			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			} else {


						log := As_logs{User:user, UserName:user.UserCode,
							OperateDatetime:timeeee,
							OperateInfo:memooo,
						}

						_, eeerr = o.Insert(&log)
						if eeerr != nil {
							beego.Error(eeerr)
							o.Rollback()
							return flag

						} else {


								errr := o.Commit()
								if errr != nil {
									beego.Error(errr)
								} else {
									flag = true
								}



				}
			}
		}

	}







	return flag

}

func Tx_SaveXuFei(keywords *As_keywords,user *As_user) bool {
	flag:=false
	o:=orm.NewOrm()
	o.Begin()
	_,errre:=o.Update(keywords,"RegPassDatetime","ServiceYears")

	timeee:=time.Now()

	if errre != nil {
		beego.Error(errre)
		o.Rollback()
		return flag

	}else {

		account:=As_account{User:user}

		account.Read("User")
		oney:=account.Money-keywords.Price
		account.Money=oney

		account.MoneyBak=account.Money
		err:=account.Update("Money","MoneyBak")









		if err != nil {
			beego.Error(err)
			o.Rollback()
			return flag

		}else {
			ServiceYears:=strconv.FormatInt(keywords.ServiceYears,10)
			Price:=strconv.FormatFloat(keywords.Price,'f',2,64)

			memo:=user.UserName+"对"+keywords.CustomName+"进行关键词申请操作,扣除预付款资金："+ServiceYears+"年"+Price+"元"
			accountdetail:=As_accountdetail{DetailType:9999,User:user,DetailTypeName:"预注册冻结资金",
				DetailDateTime:timeee,Money:keywords.Price,AccountMoney:account.Money, Memo:memo }

			eeerr:=accountdetail.Insert()






			if eeerr != nil {
				beego.Error(eeerr)
				o.Rollback()
				return flag

			}else {



				log:=As_logs{User:user,UserName:user.UserCode,
					OperateDatetime:timeee,OperateInfo:memo}
				errrer:=log.Insert()
				if errrer != nil {
					beego.Error(errrer)
					o.Rollback()
					return flag

				}else {

					errr := o.Commit()
					if errr != nil {
						beego.Error(errr)
					} else {
						flag = true
					}
				}
			}
		}


	}







	return flag

}