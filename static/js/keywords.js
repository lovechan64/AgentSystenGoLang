var customID=-1;
var customName="";
var aaa=0;
$().ready(function(){
	mover(0);
	$("#searchusertext").keyup(function(){
	 
		loadCustoms();
	});
	
	$("#searchusertext").click(function(){
		loadCustoms();
	});
	
 
	 $("#keyword").blur(function(){
				var keyword=$("#keyword").val();
			var servicetype=$("#servicetype").val();
		var years=$("#serviceyear").val();
	 if(keyword==null||keyword.length==0)
		 $("#keywordtip").html("<font color='red'>对不起,您申请的关键词不能为空,请重新填写</font>");
	 else{
		 $("#keywordtip").html("<font color='green'>正在验证关键词是否已被抢占</font>");
		 $.post("/keyword_valikey",{"keywords.keywords":keyword},function(result){
				 if(result=="success"){
					 $("#keywordtip").html("<font color='green'>恭喜您【"+keyword+"】可以被注册</font>");
					 aaa=0;
				 }else if(result=="nosuccess"){
					 $("#keywordtip").html("<font color='red'>对不起,您申请的关键词【"+keyword+"】已被抢占</font>");
					 aaa=1;
				 }

			},"html");
	 }


	 });



 $("#servicetype").change(function(){
		 jisuan();
	 });
 $("#serviceyear").change(function(){
	 jisuan();
 });

 $("#submitkeyword").click(function(){
	var servicetype= $("#servicetype").val();
	var serviceyear=$("#serviceyear").val();
	var keyword=$("#keyword").val();

	 $.post("/keyword_valikey",{"keywords.keywords":keyword},function(result){


		 if(result=="success"){
			 $("#keywordtip").html("<font color='green'>恭喜您【"+keyword+"】可以被注册</font>");
			 aaa=0;
		 }else if(result=="nosuccess"){
			 $("#keywordtip").html("<font color='red'>对不起,您申请的关键词【"+keyword+"】已被抢占</font>");
			 aaa=1;
		 }

	},"html");


	if(customID<0){
		humane.error("对不起,请选择客户");

		$("#searchusertext").focus();
	}else if($("#keyword").val()==''||$("#keyword").val().length<.0){
		$("#keywordtip").html("<font color='red'>对不起,您申请的关键词不能为空</font>");
	}else if(servicetype==""){
		humane.error("请选择服务类型");
		$("#servicetype").focus();
	}else if(aaa==1){
		humane.error("对不起,您申请的关键词【"+keyword+"】已被抢占");
		$("#keyword").focus();
	}else if(serviceyear==""){
		humane.error("请选择服务年限");
		$("#serviceyear").focus();
	}else{
		$.post("/keyword_submitkeyword",{"p":servicetype+"-"+serviceyear,"keywords.keywords":$("#keyword").val(),"keywords.customId":customID,"keywords.customName":customName},function(result){
			if(result=="exception"){
				humane.error("对不起，您提交的关键词出现异常 ，请重试");
			}else if(result=="success"){
				$.post("/keyword_account",{},function(result){
					if(result!="failed"){

						$("#accountspan").html(result);
						humane.success("恭喜您,您提交的关键词["+$("#keyword").val()+"]申请成功");
					}else{
						humane.success("对不起,您提交的关键词["+$("#keyword").val()+"]申请失败");
					}


				},'html');

			}else if(result=="nomoney"){
				humane.success("对不起,余额不足,您提交的关键词["+$("#keyword").val()+"]申请失败");
			}


		},'html');
	}


 });


});

function yanzheng(){
	var keyword=$("#keyword").val();
var servicetype=$("#servicetype").val();
var years=$("#serviceyear").val();
if(keyword==null||keyword.length==0)
$("#keywordtip").html("<font color='red'>对不起,您申请的关键词不能为空,请重新填写</font>");
else{
$("#keywordtip").html("<font color='green'>正在验证关键词是否已被抢占</font>");
$.post("/keyword_valikey",{"keywords.keywords":keyword},function(result){
	 if(result=="success"){
		 $("#keywordtip").html("<font color='green'>恭喜您【"+keyword+"】可以被注册</font>");

	 }else if(result=="nosuccess"){
		 $("#keywordtip").html("<font color='red'>对不起,您申请的关键词【"+keyword+"】已被抢占</font>");

	 }

},"html");
}


}

function jisuan(){

	var servicetype= $("#servicetype").val();
	var serviceyear=$("#serviceyear").val();
	if(servicetype==""){
		humane.error("请选择服务类型");
		$("#servicetype").focus();
	}else if(serviceyear==""){
		humane.error("请选择服务年限");
		$("#serviceyear").focus();
	}else{
		$.post("/keyword_jisuan",{"p":servicetype+"-"+serviceyear},function(result){
			if(result=="exception"){
				humane.error("计算价格出现异常,请重试");
			}else{
				$("#price").val(result);

			}


		},'html');
	}
}

function loadCustoms(){

	$.post("/keyword_searchcustom",{"custom.customName":$("#searchusertext").val()},function(result){
		var customList="<ul>";
		for(var i=0;i < result.length ; i++){
			
			customList=customList+"<li onclick=\"confirmCustom('"+result[i].Id
			+"','"+result[i].CustomName+"')\"><span>客户类型:"
			+result[i].CustomTypeName+"</span><br>"
			+result[i].CustomName+"</li>";
		}
		customList+="</ul>";
		
		$("#serachresult").html(customList);
		$("#serachresult").slideDown();
		
		
	},"json");
}
function confirmCustom(id,name){
	 customID=id;
	 customName=name;
	$("#searchusertext").val(name);
	$("#serachresult").slideUp();
	$(".customname").val(name);
}