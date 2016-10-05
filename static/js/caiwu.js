var userid;
$().ready(function(){
	userid=-1;
	mover(3);
	$("#searchUserText").keyup(function(){
		loadUser();
	});
	
	$("#searchUserText").click(function(){
		loadUser();
	});
	
	$("#caiwuok").click(function(){
		 
		if(userid<0)
			humane.error("对不起,您没有选择用户,请搜索进行用户选择");
		else if($("#zijintype").val()==''){
		humane.error("对不起,您没有选择操作类型,请选择");
		
	}else if($("#zijin").val=''||$("#zijin").val.length<=0){
		humane.error("对不起,您输入资金");
	}else{
		if(confirm("是否确认提交当前操作?")){
			$.post("/caiwu_opeaccount",{
				'account.userId':userid,
				'account.userName':$("#searchUserText").val(),
				'account.money':$("#zijin").val(),
				'accountDetail.detailType':$("#zijintype").val(),
				'accountDetail.detailTypeName':$("#zijintype option:selected").text(),
				'accountDetail.memo':$("#memo").val()



			},function(result){
				if(result=='success'){
					humane.success("当前操作成功");
					$("#systemtip").html("当前操作成功");
				}else{
					humane.error("对不起,当前操作失败");
					$("#systemtip").html("对不起,当前操作失败");
				}
			},'html');
		}
	}

	});
});
function confirmUser(uid,ucode){
	userid=uid;
	$("#searchUserText").val(ucode);

	$("#searchresult").html("");


}

function loadUser(){
	$.post("/caiwu_searchuser",{'user.userCode':$("#searchUserText").val()},
	function(result){
		var userList="<ul>";
		for(var i=0;i<result.length;i++){
			userList=userList+"<li onclick=\"confirmUser('"+result[i].Id
					+"','"+result[i].UserCode+"')\">"+result[i].UserCode
					+"</li>";
		}
		
		userList=userList+"</ul>";
		$("#searchresult").html(userList);
	},"json"		
	);
};