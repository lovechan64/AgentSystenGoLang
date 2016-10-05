$().ready(function(){
	mover(3);
	
	
	$(".deleteUser").click(function(){
		var obj=$(this);
		var	userName=obj.attr("uname");
		var	id=obj.attr("uid");
		var	roleid=obj.attr("roleid");
		if(roleid==1){
			alert("系统管理员不能被删除");
		}else{
		 if(confirm("您确定要删除【"+userName+"】吗？")) {
			 
			 var add='delete';
			 
				$.post("/userlist_edituser",{'user.id':id,'type':add}
				,function(result){

					if(result=="success"){
						humane.error('删除成功');
						window.location.href='/userlist';

					}else{
						humane.error('删除失败');

					}

				},"html");

		 }}

	});



	$("#addUser").click(function(){
		$("#addUserDiv").slideDown();
		//$("#modifyUserDiv").slideUp();

	});
	$("#cancleAdd").click(function(){
		$("#addUserDiv").slideUp();
		//$("#modifyUserDiv").slideUp();

	});

	$("#addUserSubmit").click(function(){
		var a_userCode=$("#a_userCode").val();
		var a_userName=$("#a_userName").val();
		var a_userPassword=$("#a_userPassword").val();
		var a_roleId=$("#a_roleId").val();
		var a_isStart=$("#a_isStart").val();
		if(a_userCode==''){
			humane.error('登录帐号不能为空');
		}else if(a_userName==''){
			humane.error('登录名称不能为空');
		}else if(a_userPassword==''){
			humane.error('登录密码不能为空');
		}else if(a_roleId==''){
			humane.error('请选择角色');
		}else{
			var add='add';
			$.post("/userlist_edituser",{'type':add,'user.userName':a_userName,'user.userPassword':a_userPassword,'user.roleId':a_roleId,'user.isStart':a_isStart,'user.userCode':a_userCode}
			,function(result){

				if(result=="success"){
					humane.error('添加成功');
					window.location.href='/userlist';

				}else if('repeat'==result){
					humane.error('登录帐号重复');

				}else{
					humane.error('添加失败');

				}

			},"html"
					);
		}
	});
	/*m_userName" type="text"
					name="user.userName" /><span>*</span>
				</li>
				<li>登录密码:<input id=""
					type="password" name="user.userPassword" /><span>*</span>
				</li>
				<li>角色:<s:select id="" list="roleList" headerKey="" headerValue="请选择"
						name="user.roleId" listKey="id" listValue="roleName"></s:select><span>*</span>
				</li>
				<li>是否启用:<select id="*/
	$(".modifyUser").click(function(){
		var obj=$(this);
		$("#m_userName").val(obj.attr("uname"));
		$("#userid").val(obj.attr("uid"));

		$("#m_roleId").val(obj.attr("roleid"));
		$("#m_isStart").val(obj.attr("isstart"));
		$("#m_userCode").val(obj.attr("ucode"));


			$("#modifyUserDiv").slideDown();
		//$("#modifyUserDiv").slideUp();

	});
	$("#cancleModify").click(function(){
		$("#modifyUserDiv").slideUp();
		//$("#modifyUserDiv").slideUp();

	});

	$("#modifyUserSubmit").click(function(){

		var	userName=$("#m_userName").val();
		var	id=$("#userid").val();
		var	userPassword=$("#m_userPassword").val();
		var	roleId=$("#m_roleId").val();
		var	isStart=$("#m_isStart").val();
		var	userCode=$("#m_userCode").val();


		if(userCode==''){
			humane.error('登录帐号不能为空');
		}else if(userName==''){
			humane.error('登录帐号不能为空');
		}else if(userPassword==''){
			humane.error('登录帐号不能为空');
		}else if(roleId==''){
			humane.error('登录帐号不能为空');
		}else{
			var add='modify';

			$.post("/userlist_edituser",{'user.id':id,'type':add,'user.userName':userName,'user.userPassword':userPassword,'user.roleId':roleId,'user.isStart':isStart,'user.userCode':userCode}
			,function(result){
				
				if(result=="success"){
					humane.error('修改成功');
					window.location.href='/userlist';
					
				}else{
					humane.error('修改失败');
					
				}
				
			},"html");
		}
	});
	
});

