$().ready(function(){
	$("#cball").change(function(){
		//var checkList='';
		var cblist=$(".cb");
		//var roleId=$("#roleid");
		for(var i=0;i<cblist.length;i++){
			cblist[i].checked=$(this).attr("checked")=='checked'?'checked':'';
		}
	});
	$("#saverolefunc").click(function(){
		var checkList='';
		var cblist=$(".cb");
		var roleId=$("#roleid").val();
		
		for(var i=0;i<cblist.length;i++){
			if(cblist[i].checked){
				checkList+=cblist[i].value+",";
			};
		}
		$.post("/userlist_saverolefunc.action",{"checkList":checkList,"roleId":roleId},function(result){
			if(result=="success"){
				humane.success("保存成功");
			}else{
				humane.success("保存失败");
			}
		},"html");
	});
	
});