Date.prototype.format=function(format){
	var o={
		"M+":this.getMonth()+1,
		"d+":this.getDate(),
		"y+":this.getFullYear()
			
			
	};
	for(var k in o){
		if(new RegExp("("+k+")").test(format))
		format=format.replace(RegExp.$1,o[k]);
	}
	return format;

}

function mover(object){
	for(var i=0;i<5;i++){
		if(i==object){
		$("#m_"+i).attr("class","m_li_a");
		$("#s_"+i).attr("class","s_li_a");
		}else{
			$("#m_"+i).attr("class","m_li");
			$("#s_"+i).attr("class","s_li");
			
		}
	}
}