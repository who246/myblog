{{template "common/head.tpl"}}
{{template "common/nav.tpl" .}}
{{template "common/roll.tpl" .}}

<div class=" ">
    <div class="col-md-8 article-list">
	
	{{range .articles}}
	
      <h2>
	<a href="/article/view/{{.Id}}.html">
        {{.Title}}
	</a>
      </h2>
	
	 <h6 class="text-muted">文章发表于 {{dateformat .CreateTime "2006-01-02 15:04:00"}}，修改于 {{dateformat .ModifyTime  "2006-01-02 15:04:00"}}</h6>
      <p>
		{{substr (html2str .Content) 0 200}}
      </p>
      
     {{end}}
	
      {{template "common/page.tpl" .}}
    </div>
     {{template "common/left.tpl" .}}
  </div>
     {{template "common/foot.tpl"}}