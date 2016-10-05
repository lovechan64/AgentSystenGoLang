function validateLoginUserFunc(){
	var flag=false;
	var usercode=$.trim($("#usercode").val());
	var userpassword=$.trim($("#password").val());
	if(usercode==""||usercode==null){
		
		$("#usercode").focus();
		alert("账号不能为空!");
	}else if(userpassword==""||userpassword==null){
		
		
		$("#password").focus();
		alert("密码不能为空!");
		
	}else{
		 

		$.post("/login",{'user.userCode':usercode,'user.Password':userpassword,"yanzheng":"1"}
		,function(data){
			if("noexitusercode"==data ){
				alert("对不起，账号不存在");
			}else if("errorpwd"==data ){
				alert("对不起，密码错误");
			}else if("failed"==data ){
				alert("对不起，系统错误");
			}else if("success"==data ){
				 flag=true;
			}
		},"html");
		
		
		
		
		
		
		
		

	}
	
	
	return flag;
	
}


$().ready(function(){
	$("#denglu").click(function(){
	 

		var flag=false;
		var usercode=$.trim($("#usercode").val());
		var userpassword=$.trim($("#password").val());
		if(usercode==""||usercode==null){
			
			$("#usercode").focus();
			alert("账号不能为空!");
		}else if(userpassword==""||userpassword==null){
			
			
			$("#password").focus();
			alert("密码不能为空!");
			
		}else{
			 

			$.post("/login",{'user.userCode':usercode,'user.userPassword':userpassword,"yanzheng":"1"}
			,function(data){
				if("noexitusercode"==data ){
					alert("对不起，账号不存在");
				}else if("errorpwd"==data ){
					alert("对不起，密码错误");
				}else if("failed"==data ){
					alert("对不起，系统错误");
				}else if("success"==data ){
					 flag=true;
					 //在js中  提交表单 用  form的id
					 $("#tijiao").submit();
				}
			},"html");
			
			
			
			
			
			
			
			

		}
		
		 
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
	});
});