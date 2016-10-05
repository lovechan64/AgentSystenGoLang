 
$().ready(function(){
	mover(1);

	
	
	$("#baocun").click(function(){
		var appUserName=$("#appUserName").val();
		var appPassword=$("#appPassword").val();
		var iosDownloadUrl=$("#iosDownloadUrl").val();
		var androidDownloadUrl=$("#androidDownloadUrl").val();
		if(appUserName==""){
			humane.error("登录名称不能为空！");
			$("#appUserName").foucs();
		}else if (appPassword==""){
			humane.error("登录密码不能为空！");
			$("#appPassword").foucs();
		}else if (iosDownloadUrl==""){
			humane.error("IOS下载地址不能为空！");
			$("#iosDownloadUrl").foucs();
		}else if (androidDownloadUrl==""){
			humane.error("安卓下载不能为空！");
			$("#androidDownloadUrl").foucs();
		}else{
			$("#kform").submit();
		}
		
		 
	 	
	});
});

 