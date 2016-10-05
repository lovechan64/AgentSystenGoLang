$().ready(function(){
	mover(3);
	
	$(".checkselect").change(function(){
		var obj=$(this);
		var v=obj.val();
		var kid=obj.attr("kid");
		var keyword=obj.attr("keyword");
		if(v==1){
			if(confirm("您确定要将"+keyword+"的状态修改为[审核中]吗？")){
				$.post("keyword_updatekeywords",{"keywords.id":kid,"keywords.checkStatus":1},function(result){

					if("success"==result){
						humane.success("恭喜您,的状态被修改为[审核中]，操作成功");
						 window.location.reload(true);

					}else{
						humane.error("对不起,的状态修改[审核中]操作失败");


					}


				},'html');

			}

		}else if(v==2){


			if(confirm("您确定要将"+keyword+"的状态修改为[审核通过]吗？")){
				$.post("keyword_updatekeywords",{"keywords.id":kid,"keywords.checkStatus":2},function(result){

					if("success"==result){
						humane.success("恭喜您,的状态被修改为[审核通过]，操作成功");
						window.location.reload(true);
					}else{
						humane.error("对不起,的状态修改[审核通过]操作失败");


					}


				},'html');

			}

		}else if(v==3){

			if(confirm("您确定要将"+keyword+"的状态修改为[审核不通过]吗？")){
				$.post("keyword_updatekeywords",{"keywords.id":kid,"keywords.checkStatus":3},function(result){

					if("success"==result){
						humane.success("恭喜您,的状态被修改为[审核不通过]，操作成功");
						window.location.reload(true);
					}else{
						humane.error("对不起,的状态修改[审核不通过]操作失败");


					}


				},'html');

			}


		}else if(v==5){


			if(confirm("您确定要将"+keyword+"的状态修改为[不使用]吗？")){
				$.post("keyword_updatekeywords",{"keywords.id":kid,"keywords.isUse":0},function(result){

					if("success"==result){
						humane.success("恭喜您,的状态被修改为[不使用]，操作成功");
						window.location.reload(true);
					}else{
						humane.error("对不起,的状态修改[不使用]操作失败");


					}


				},'html');

			}


		}else if(v==6){

			if(confirm("您确定要将"+keyword+"的状态修改为[使用]吗？")){
				$.post("keyword_updatekeywords",{"keywords.id":kid,"keywords.isUse":1},function(result){
					
					if("success"==result){
						humane.success("恭喜您,的状态被修改为[使用]，操作成功");
						window.location.reload(true);
					}else{
						humane.error("对不起,的状态修改[使用]操作失败");
						
						
					}
					
					
				},'html');
				
			}
		}
	});
	
	
});