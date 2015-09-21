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
             
			<div id="cyReping" role="cylabs" data-use="reping"></div>
			
			<script type="text/javascript" charset="utf-8" src="http://changyan.itc.cn/js/??lib/jquery.js,changyan.labs.js?appid=cyrYGvclA"></script>
		</div>            
			 
		</div>