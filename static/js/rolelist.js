$().ready(function(){
	mover(3);
	$("#addRole").click(function(){
		
		$("#addRolediv").slideDown();
		
		
	});
	
	
	/*a_roleName" type="text" name="role.roleName"/></li>

			<li>是否启用:<select id="a_isStart" name="role.isStart">
				<option value="1" selected="selected">启用</option>
				<option value="0">停用</option>
			
			</select>
			
			
			<input type="button" id="addRoleBtn" value="保存"/>
			<input type="button" id="cancleAddRoleBtn*/
	$("#addRoleBtn").click(function(){
		 
		var roleName= $("#a_roleName").val();
		var isStart= $("#a_isStart").val();
		var add="add";
	
		
		if(roleName==''){
			humane.error("角色名称不能为空");
		}else{
			$.post("/rolelist_editrole.action",{"type":add,"role.roleName":roleName, "role.isStart":isStart }
			,function(result){
				if(result=="peat"){
					humane.error("对不起，该角色名称已存在");
						}else{
							
							humane.success("已添加成功");
							$("#addRolediv").slideUp();
					window.location.reload(true);		
						}
				
			},"html");
			
		}
		
		
	});
	$("#cancleAddRoleBtn").click(function(){
		
		$("#addRolediv").slideUp();
		
		
		
	});
	
	
	
	
	$(".modifyRole").click(function(){
		
		/*isstart="<s:property value="isStart"  />"
			rolename="<s:property value="roleName"  />"
			
			
			 roleid="<s:property value="id" />">修改</span>*/
		
		var obj= $(this);
		$("#roleid").val(obj.attr("roleid"));
		$("#m_roleName").val(obj.attr("rolename"));
		$("#m_isStart").val(obj.attr("isstart"));

			$("#modifyRoleDiv").slideDown();
		
		
	});
	
	$("#modifyCancle").click(function(){
		
		$("#modifyRoleDiv").slideUp();
		
		
		
	});
	$("#modifyRoleSubmit").click(function(){
		
	 var id=	$("#roleid").val();
		 
		var roleName= $("#m_roleName").val();
		var isStart= $("#m_isStart").val();
		 var type="modify";
		
		
		if(roleName==''){
			humane.error("角色名称不能为空");
		}else{
			$.post("/rolelist_editrole.action",{"type":type,"role.roleName":roleName, "role.isStart":isStart ,"role.id":id}
			,function(result){
				 
							
							humane.success("已修改成功");
							$("#modifyRoleDiv").slideUp();
					window.location.reload(true);		
						 
				
			},"html");
			
		}
		
		
	});
	
	

	
	
	
	$(".deleteRole").click(function(){
		var obj= $(this);
		$("#roleid").val(obj.attr("roleid"));
		$("#deleteRoleName").val(obj.attr("rolename"));
			$("#deleteRolediv").slideDown();
		
		
	});
	
	$("#cancleDeleteRoleBtn").click(function(){
		
		$("#deleteRolediv").slideUp();
		
		
		
	});
	$("#deleteRoleBtn").click(function(){
		 
	 var id=	$("#roleid").val();
		 
			$.post("/rolelist_deleterole.action",{"role.id":id}
			,function(result){
				 
							
							humane.success("已删除成功");
							$("#deleteRolediv").slideUp();
					window.location.reload(true);		
						 
				
			},"html");
			
	 
		
		
	});
	/*simpleconfig">
		
		<h3>服务年限管理-填写最大的服务年限</h3>
		<input value="<s:property value="role.id"  />" type="hidden" id="simpleId">
		<input type="hidden" id="configType" value="<s:property value="configType"  />">
		<ul>
			<li>配置名称:<input value="<s:property value="role.configTypeName"  />" type="text" id="simpleTypeName"></li>
			<li>配置数值:<input value="<s:property value="role.configValue"  />" type="text" id="simpleConfigValue"></li>
			<li><input value="保存" type="button" id="simpleBtn"></li>
		</ul>
	</div>*/
	
/*	$("#simpleBtn").click(function(){
		var configType=	$("#configType").val();
		var simpleConfigId=	$("#simpleConfigId").val();
		var simpleTypeName=	$("#simpleTypeName").val();
		var simpleConfigValue=	$("#simpleConfigValue").val();
			if(simpleTypeName==''){
				humane.error("配置名称不能为空");
				
			}else if(simpleConfigValue==""){
				
				humane.error("配置数值不能为空");
				
			}else{
				
				$.post("/modifyconfig.action",{"role.configTypeName":simpleTypeName,"role.id":simpleConfigId,"role.configValue":simpleConfigValue}
				,function(result){
					 
								
								humane.success("已修改成功");
								$("#modifyRolediv").slideUp();
						window.location.reload(true);		
							 
					
				},"html");
				
			}
	});*/
});

