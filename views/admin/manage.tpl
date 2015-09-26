<script type="text/javascript">
$(document).ready(function(){
  
	$(".del").click(function(){
     var url = $(this).attr("url");
	 var link = $(this).attr("link");
	 var ids = $("#"+link).val();
	if (confirm("是否确认！") == false) return;   
	$.post(url,{ids:ids},function(r){
      alert(r.msg);
     }, "json");
   });


$(".mody").click(function(){
     var url = $(this).attr("url");
	 var link = $(this).attr("link");
	 var ids = $("#"+link).val();
     window.location.href=url+"?id="+ids;
});
 
  

   $(".banner").click(function(){
     var id = $("#id").val();
	 var url = $("#url").val();
	 var n = $("#n").val(); 
	$.post("/admin/config/set",{id:id,url:url,n:n},function(r){
      alert(r.msg);
     }, "json");
   });
});
</script>
<div class="row">
  
  <div class="col-lg-6">
    <div class="input-group">
      <input type="text" class="form-control" id="aids" placeholder="输入文章Id例如1,2,3,4；如果更新输入一个id">
      <span class="input-group-btn">
        <button class="btn btn-default del" url="/admin/aticle/del" link="aids" type="button">删除</button>
      </span>
	  <span class="input-group-btn">
        <button class="btn btn-default mody" url="/admin/aticle/showmodify" link="aids" type="button">更新</button>
      </span>
    </div><!-- /input-group -->
  </div><!-- /.col-lg-6 -->
</div><!-- /.row -->
<hr/>
<div class="row">
  
  <div class="col-lg-6">
    <div class="input-group">
      <input type="text" class="form-control" id="tids" placeholder="输入类型Id例如1,2,3,4">
      <span class="input-group-btn">
       <button class="btn btn-default del" url="/admin/type/del" link="tids" type="button">删除</button>
      </span>
    </div><!-- /input-group -->
  </div><!-- /.col-lg-6 -->
</div><!-- /.row -->
 <hr/>
<div class="row">
  
  <div class="col-lg-6">
    <div class="input-group">
	  <span>
      <input type="text" class="form-control" id="url" placeholder="输入url">
	  <input type="text" class="form-control" id="id" placeholder="输入文章id">
	  <input type="text" class="form-control" id="n" placeholder="输入bannerid">
	 <span>
      <span class="input-group-btn">
       <button class="btn btn-default banner"  type="button">添加</button>
      </span>
    </div><!-- /input-group -->
  </div><!-- /.col-lg-6 -->
</div><!-- /.row -->

<hr/>
<div class="row">
  
  <div class="col-lg-6">
    <div class="input-group"> 
      <span class="input-group-btn">
       <button class="btn btn-default del" url="/admin/construct/update" link="tids" type="button">更新并重启系统</button>
      </span>
    </div><!-- /input-group -->
  </div><!-- /.col-lg-6 -->
</div><!-- /.row -->
<hr/>
<div class="row">
  
  <div class="col-lg-6">
    <div class="input-group"> 
      <span class="input-group-btn">
       <button class="btn btn-default del" url="/admin/construct/bulid" link="tids" type="button">重新构建并重启系统（需要配置环境）</button>
      </span>
    </div><!-- /input-group -->
  </div><!-- /.col-lg-6 -->
</div><!-- /.row -->