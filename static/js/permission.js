$().ready(function(){
	mover(3);
	/*$("[name='functioncheck']").each(function(){
		if($(this).attr("checked")==true){
			 var id=	$(this).val();
			 var isStart=1;
				$.post("/updaterolepremission.action",{"rolePremission.id":id,"rolePremission.isStart":isStart}
				,function(result){},"html");
		}else{
			 var id=	$(this).val();
			 var isStart=0;
				$.post("/updaterolepremission.action",{"rolePremission.id":id,"rolePremission.isStart":isStart}
				,function(result){},"html");
			
			
		}
	});*/
	/*
	
	
	$("#addsystemconfig").click(function(){
		
		$("#addSystemdiv").slideDown();
		
		
	});
	$("#addSystemConfigBtn").click(function(){
		var configType= $("#addConfigType").val();
		var configTypeName= $("#addConfigName").val();
		var isStart= $("#addIsStartSelect").val();
	
		var configValue;
		if($("#addConfigValue").val()!=null){
			configValue=$("#addConfigValue").val();
		}
		
		
		if(configType==''){
			humane.error("配置类型不能为空");
		}else if(configTypeName==''){
			humane.error("配置类型名称不能为空");
			
		}else{
			$.post("/addconfig.action",{"systemConfig.configType":configType,"systemConfig.configTypeName":configTypeName,"systemConfig.isStart":isStart,"systemConfig.configValue":configValue}
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
		$("#modifyIsStartSelect").val(obj.attr("isstart"));
		if(obj.attr("configvalue")!=null){
		$("#modifyConfigValue").val(obj.attr("configvalue"));
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
		
		
		if(configType==''){
			humane.error("配置类型不能为空");
		}else if(configTypeName==''){
			humane.error("配置类型名称不能为空");
			
		}else{
			$.post("/modifyconfig.action",{"systemConfig.configTypeName":configTypeName,"systemConfig.isStart":isStart,"systemConfig.id":id,"systemConfig.configValue":configValue}
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
		 
			$.post("/deleteconfig.action",{"systemConfig.id":id}
			,function(result){
				 
							
							humane.success("已删除成功");
							$("#deleteSystemdiv").slideUp();
					window.location.reload(true);		
						 
				
			},"html");
			
	 
		
		
	});
	simpleconfig">
		
		<h3>服务年限管理-填写最大的服务年限</h3>
		<input value="<s:property value="systemConfig.id"  />" type="hidden" id="simpleId">
		<input type="hidden" id="configType" value="<s:property value="configType"  />">
		<ul>
			<li>配置名称:<input value="<s:property value="systemConfig.configTypeName"  />" type="text" id="simpleTypeName"></li>
			<li>配置数值:<input value="<s:property value="systemConfig.configValue"  />" type="text" id="simpleConfigValue"></li>
			<li><input value="保存" type="button" id="simpleBtn"></li>
		</ul>
	</div>
	
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
				
				$.post("/modifyconfig.action",{"systemConfig.configTypeName":simpleTypeName,"systemConfig.id":simpleConfigId,"systemConfig.configValue":simpleConfigValue}
				,function(result){
					 
								
								humane.success("已修改成功");
								$("#modifySystemdiv").slideUp();
						window.location.reload(true);		
							 
					
				},"html");
				
			}
	});*/
});

