$().ready(function(){

	mover(0);



	$("#button").click(function(){
		var keywordsid=$("#keywordsid").val();
        		var appPassword=$("#appPassword").val();
                		var appUserName=$("#appUserName").val();
		if(appPassword==""){
			hamane.error("对不起，请输入APP密码");
			$("#appPassword").focus();
		}else if(appUserName==""){
			hamane.error("对不起，请输入APP名称");
			$("#appUserName").focus();
		}else{
			$.post("/keywordmanage_modifyapp",{
				'keywords.id':keywordsid,
                				'keywords.appUserName':appUserName,
                                				'keywords.appPassword':appPassword},function(result){
				if(result=='success'){
					humane.success("恭喜您，开通APP成功！");

				}else  {
					humane.success("对不起，，开通APP失败！");

				}
			},'html');
		}
	});
});

