function serchReportFunc(){
	/*if(comtime($("#starttime").val(),$("#endtime").val()))
	{
		humane.error("开始时间大于结束时间，请检查");
		return false;
	}*/
	return true;
}

/*function comtime(s,e){
	var a=s.split("-");
	var sst=NEW Date(a[0],a[1],a[2]);
	var ss=sst.getTime();
	
	
	var a1=s.split("-");
	var sst1=NEW Date(a1[0],a1[1],a1[2]);
	var ss2=sst1.getTime();
	
	return ss>ss2;
	
	
}*/
 

$().ready(function(){
	mover(2);
}); 