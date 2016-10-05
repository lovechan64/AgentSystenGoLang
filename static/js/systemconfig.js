$().ready(function(){
	mover(4);
	$("#addsystemconfig").click(function(){
		
		$("#addSystemdiv").slideDown();
		
		
	});
	$("#addSystemConfigBtn").click(function(){
		 
		var configType= $("#addConfigType").val();
		var configTypeName= $("#addConfigName").val();
		var isStart= $("#addIsStartSelect").val();
		var configTypeName1= $("#addConfigName1").val();
		var configTypeName2= $("#addConfigName2").val();
	
		
		
		var configValue;
		if($("#addConfigValue").val()!=null){
			configValue=$("#addConfigValue").val();
		}
		if(configTypeName1!=''&&configTypeName1!=null&&configTypeName2!=''&&configTypeName2!=null){
			configTypeName="买"+configTypeName1+"赠"+configTypeName2;
			configValue=configTypeName1;
			
		}
		 
		if(configType==''){
			humane.error("配置类型不能为空");
		}else if(configTypeName==''){
			humane.error("配置类型名称不能为空");
			
		}else{
			
			$.post("/servicetype_addconfig",{"systemConfig.configType":configType,"systemConfig.configTypeName":configTypeName,"systemConfig.isStart":isStart,"systemConfig.configValue":configValue}
			,function(result){
				if(result=="peat"){
					humane.error("对不起，该类型名称已存在");
						}else{

							humane.success("已添加成功");
							$("#addSystemdiv").slideUp();
					window.location.reload(true);
						}

			},"html");

		}


	});
	$("#cancleAddSystemConfigBtn").click(function(){

		$("#addSystemdiv").slideUp();



	});




	$(".modifySystemBtn").click(function(){
		var obj= $(this);
		$("#sid").val(obj.attr("sid"));
		$("#modifyConfigName").val(obj.attr("sname"));
		$("#configtype").val(obj.attr("configtype"));

		$("#modifyIsStartSelect").val(obj.attr("isstart"));
		if(obj.attr("configvalue")!=null){
		$("#modifyConfigValue").val(obj.attr("configvalue"));
		}
		if(obj.attr("configtype")==7){
			var a= obj.attr("sname").split("买");
			var b=a[1];
			var c=b.split("赠");
			$("#modifyConfigName1").val(c[0]);
			$("#modifyConfigName2").val(c[1]);



		}







			$("#modifySystemdiv").slideDown();


	});

	$("#cancleModifySystemConfigBtn").click(function(){

		$("#modifySystemdiv").slideUp();



	});
	$("#modifySystemConfigBtn").click(function(){

	 var id=	$("#sid").val();
		var configType= $("#modifyConfigType").val();
		var configTypeName= $("#modifyConfigName").val();
		var isStart= $("#modifyIsStartSelect").val();
		var configValue;
		if($("#modifyConfigValue").val()!=null){
			configValue=$("#modifyConfigValue").val();
		}
		if($("#configtype").val()==7){

			configTypeName="买"+$("#modifyConfigName1").val()+"赠"+$("#modifyConfigName2").val();
			configValue=$("#modifyConfigName1").val();

		}


		if(configType==''){
			humane.error("配置类型不能为空");
		}else if(configTypeName==''){
			humane.error("配置类型名称不能为空");

		}else{
			$.post("/servicetype_modifyconfig",{"systemConfig.configTypeName":configTypeName,"systemConfig.isStart":isStart,"systemConfig.id":id,"systemConfig.configValue":configValue}
			,function(result){


							humane.success("已修改成功");
							$("#modifySystemdiv").slideUp();
					window.location.reload(true);


			},"html");

		}


	});






	$(".deleteSystemBtn").click(function(){
		var obj= $(this);
		$("#deleteid").val(obj.attr("sid"));
		$("#deleteConfigName").val(obj.attr("sname"));

			$("#deleteSystemdiv").slideDown();


	});

	$("#cancleDeleteSystemConfigBtn").click(function(){

		$("#deleteSystemdiv").slideUp();



	});
	$("#deleteSystemConfigBtn").click(function(){

	 var id=	$("#deleteid").val();

			$.post("/servicetype_deleteconfig",{"systemConfig.id":id}
			,function(result){


							humane.success("已删除成功");
							$("#deleteSystemdiv").slideUp();
					window.location.reload(true);


			},"html");




	});
	/*simpleconfig">

		<h3>服务年限管理-填写最大的服务年限</h3>
		<input value="<s:property value="systemConfig.id"  />" type="hidden" id="simpleId">
		<input type="hidden" id="configType" value="<s:property value="configType"  />">
		<ul>
			<li>配置名称:<input value="<s:property value="systemConfig.configTypeName"  />" type="text" id="simpleTypeName"></li>
			<li>配置数值:<input value="<s:property value="systemConfig.configValue"  />" type="text" id="simpleConfigValue"></li>
			<li><input value="保存" type="button" id="simpleBtn"></li>
		</ul>
	</div>*/

	$("#simpleBtn").click(function(){
		var configType=	$("#configType").val();
		var simpleConfigId=	$("#simpleConfigId").val();
		var simpleTypeName=	$("#simpleTypeName").val();
		var simpleConfigValue=	$("#simpleConfigValue").val();
			if(simpleTypeName==''){
				humane.error("配置名称不能为空");

			}else if(simpleConfigValue==""){

				humane.error("配置数值不能为空");

			}else{

				$.post("/servicetype_modifyconfig",{"systemConfig.configTypeName":simpleTypeName,"systemConfig.id":simpleConfigId,"systemConfig.configValue":simpleConfigValue}
				,function(result){
					 
								
								humane.success("已修改成功");
								$("#modifySystemdiv").slideUp();
						window.location.reload(true);		
							 
					
				},"html");
				
			}
	});
});

