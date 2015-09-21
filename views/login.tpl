{{template "common/head.tpl" .}}
{{template "common/nav.tpl" .}}
{{template "common/roll.tpl" .}}
<div class="container-fluid">
 
	 
	<div class="article-list">
		<div class="col-md-6">
			<div class="alert alert-dismissable alert-info">
				 
				<button type="button" class="close" data-dismiss="alert" aria-hidden="true">
					
				</button>
				<h4>
					后台管理登入
				</h4> <strong>闲人莫入!</strong> 误入请点击<a href="/" class="alert-link">这里</a>返回首页
			</div>
			{{if .error}}
			<div class="alert alert-dismissable alert-danger">
				 
				<button type="button" class="close" data-dismiss="alert" aria-hidden="true">
					×
				</button>
				<h4>
					错误!
				</h4> {{.error}}
			</div>
			{{end}}
		</div>
		
		<div class="col-md-6">
			<form class="form-horizontal" role="form" method="POST" action="/login" >
				<div class="form-group">
					 
					<label for="inputEmail3" class="col-sm-2 control-label">
						用户名
					</label>
					<div class="col-sm-10">
						<input type="username" name="username" required class="form-control"  pattern="^[a-zA-Z0-9_]{5,12}$"  title="用户名大于4位小于12位"  value="{{.username}}"/>
					</div>
				</div>
				<div class="form-group">
					 
					<label for="inputPassword3" class="col-sm-2 control-label">
						密码
					</label>
					<div class="col-sm-10">
						<input type="password" name="password" required class="form-control" id="inputPassword3" value="{{.password}}" />
					</div>
				</div>
				<div class="form-group">
					<div class="col-sm-offset-2 col-sm-10">
						<div class="checkbox">
							 
							<label>
								<input type="checkbox" name="isRemeber"/> 记住
							</label>
						</div>
					</div>
				</div>
				<div class="form-group">
					<div class="col-sm-offset-2 col-sm-10">
						 
						<button type="submit" class="btn btn-default">
							登入
						</button>
					</div>
				</div>
			</form>
		</div>
	</div>
</div>
{{template "common/foot.tpl"}}