$().ready(function(){
	mover(0);
	$("#modifypwdbtna").click(function(){
		
		$("#modifydiv").show();
		$("#oldpwdtex").val('');
		$("#newpwd").val('');
		$("#newpwd2").val('');
		
	});
	
	$("#oldpwdtext").blur(function(){
		if($("#oldpwdtext").val()!=""){
			
			$("#oldpwdtexttip").css("color","green");
		}else{
			
			$("#oldpwdtexttip").css("color","red");
			
		}
		
		
		
	});
	$("#newpwd").blur(function(){
		
		if($("#newpwd").val()!=""){
			
			$("#newpwdtip").css("color","green");
		}else{
			
			$("#newpwdtip").css("color","red");
			
		}
		
	});
	$("#newpwd2").blur(function(){
		
		
		if($("#newpwd2").val()!=""){
			
			$("#newpwd2tip").css("color","green");
		}else{
			
			$("#newpwd2tip").css("color","red");
			
		}
		
	});
	
	
	$("#modifypwdbtn").click(function(){
	 var oldpwd=	$("#oldpwdtext").val();
	 var newpwd=	$("#newpwd").val();
	 var newpwd2=	$("#newpwd2").val();
		if(oldpwd.length<6){
			$("#oldpwdtext").focus();
			humane.error("请输入原始密码");
			
			
		}else if(newpwd.length<6){
			$("#newpwd").focus();
			humane.error("密码不能小于6位");
		}else if(newpwd2.length<6){
			$("#newpwd2").focus();
			humane.error("密码不能小于6位");
		}else if(newpwd!=newpwd2){
			humane.error("两次密码不相同");
			
		}else{
			$.post("/login_modifypwd",{"user.userName":oldpwd,'user.userPassword':newpwd},function(result){
				if(result=="success"){
					humane.success("修改密码成功")
					$("#modifydiv").hide();
					
				}else{
					humane.error("密码修改失败");
					
					
				}
			},'html');
			
		}
	});
	$("#modifypwdcancel").click(function(){
		
		$("#modifydiv").hide();
		$("#oldpwdtex").val('');
		$("#newpwd").val('');
		$("#newpwd2").val('');
		
		
		
	});
	 
})