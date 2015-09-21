<script>
	$(window).load(function () {
            $("pre").addClass("linenums");
            prettyPrint();
     }) 
			var editor;
			KindEditor.ready(function(K) {
				editor = K.create('textarea[name="content"]', {
					allowFileManager : true
				});
				K('input[name=getHtml]').click(function(e) {
					alert(editor.html());
				});
				K('input[name=isEmpty]').click(function(e) {
					alert(editor.isEmpty());
				});
				K('input[name=getText]').click(function(e) {
					alert(editor.text());
				});
				K('input[name=selectedHtml]').click(function(e) {
					alert(editor.selectedHtml());
				});
				K('input[name=setHtml]').click(function(e) {
					editor.html('<h3>Hello KindEditor</h3>');
				});
				K('input[name=setText]').click(function(e) {
					editor.text('<h3>Hello KindEditor</h3>');
				});
				K('input[name=insertHtml]').click(function(e) {
					editor.insertHtml('<strong>插入HTML</strong>');
				});
				K('input[name=appendHtml]').click(function(e) {
					editor.appendHtml('<strong>添加HTML</strong>');
				});
				K('input[name=clear]').click(function(e) {
					editor.html('');
				});
			});
		</script>
	 	<h3>编辑器</h3>
	<form method="POST" action="{{.op}}">
			{{range .err}}
	<div style="color:#F00"  class="col-md-8">
	{{.Key}}{{.Message}}
     </div>
	{{end}}
	   <label for="title">类型</label>
	    <input type="hidden"  name="id" value="{{.art.Id}}" />
	    <select  class="form-control"  name="typeId">
		<option value="－1" >请选择</option>
	
		{{range .types}} 
		 <option value="{{.Id}}"  {{if eq .Id $.art.TypeId}} selected="selected" {{end}} >{{.Name}}</option>
		 
		{{end}}
		</select>
	   <label for="title">标题</label>
	    <input type="title" class="form-control" id="title" name="title" placeholder="标题" value="{{.art.Title}}"/>

			<textarea name="content" style="width:800px;height:400px;visibility:hidden;" placeholder="Enter text ..." >{{.art.Content}}</textarea>
			<input class="btn btn-danger"  type="submit" value="Submit">
			 
		</form>