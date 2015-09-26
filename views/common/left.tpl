	<div class="col-md-4">
			<div class="page-header">
				<h1>
					 <small>欢迎来到我的博客</small>
				</h1>
			</div>
 
			 <div class="list-group">
                    <a href="#" class="list-group-item disabled">分类</a> 
					{{range .types}}
                    <a href="/?tid={{.Id}}" class="list-group-item"><span class="badge">{{.Nums}}</span>{{.Name}}</a> 
					{{end}}
                </div>
				
		<div class="page-header">
             
		<!-- 代码1：放在页面需要展示的位置  -->
<!-- 如果您配置过sourceid，建议在div标签中配置sourceid、cid(分类id)，没有请忽略  -->
<div id="cyReping" role="cylabs" data-use="reping"></div>
<!-- 代码2：用来读取评论框配置，此代码需放置在代码1之后。 -->
<!-- 如果当前页面有评论框，代码2请勿放置在评论框代码之前。 -->
<!-- 如果页面同时使用多个实验室项目，以下代码只需要引入一次，只配置上面的div标签即可 -->
<script type="text/javascript" charset="utf-8" src="http://changyan.itc.cn/js/??lib/jquery.js,changyan.labs.js?appid=cyrYGvclA"></script>		</div>            
			 
		</div>