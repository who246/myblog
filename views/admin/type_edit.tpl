<script type="text/javascript">
$(document).ready(function(){
  
	$(".add").click(function(){
     var url = $(this).attr("url");
	 var link = $(this).attr("link");
	 var name = $("#"+link).val();  
	$.post(url,{name:name},function(r){
      alert(r.msg);
     }, "json");
   });

});
 
</script>
<div class="row">
  
  <div class="col-lg-6">
    <div class="input-group">
      <input type="text" class="form-control" id="name" placeholder="类型名称">
      <span class="input-group-btn">
        <button class="btn btn-default add" url="/admin/type/add" link="name" type="button">添加</button>
      </span>
	 
    </div><!-- /input-group -->
  </div><!-- /.col-lg-6 -->
</div>