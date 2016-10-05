$().ready(function(){
	 
	mover(0);





	
	$(".delectkeyword").click(function(){
		var obj=$(this);
		var kid=obj.attr("kid");
		var keyword=obj.attr("keyword");
		
		
		
		if(confirm("为了数据安全,需要您两次确认才能删除，您确定要删除关键词["+keyword+"]吗?(一次)"))
			if(confirm("为了数据安全,需要您两次确认才能删除，您确定要删除关键词["+keyword+"]吗?(两次)")){
				
				$.post("/keywordmanage_delectkeyword",{
					 
					'keywords.id':kid},function(result){
					if(result=='success'){
						humane.success("删除成功！");
						 window.location.reload(true);
					}else  {
						humane.success("删除成功！");
					
					}
				},'html');
				
				
			}
	 
		
		
		
	});
});
function callback(){
	window.location.reload(true);
}
