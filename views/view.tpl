{{template "common/head.tpl" .}}
{{template "common/nav.tpl"}}

<script>
	$(window).load(function () {
            $("pre").addClass("linenums");
            prettyPrint();
     }) 
</script>

<div class=" ">
    <div class="col-md-8 article-list">
	<header class="page-header">
    
      <h1 class="entry-title">{{.article.Title}}</h1>
      <p class="meta">
        
      <time data-updated="true" pubdate="">{{dateformat .article.ModifyTime "2006-01-02 15:04:00"}} </time>
        
		
         | <a href="#SOHUCS">评论</a> | 所属分类:{{range .types}} {{if eq .Id $.article.TypeId}} <a href="/?tid={{.Id}}">{{.Name}}</a> {{end}} {{end}} | 浏览次数:{{.article.ClickNum}} 
        
      </p>
    
  </header>
	 <div class="entry-content">
     {{str2html .article.Content}}

      </div>
	<div style="padding-top:20px">         
            <p style="font-size:12px;">版权声明：本文为博主原创文章，未经博主允许不得转载。</p>
    </div>
    {{template "common/comment.tpl" .}}
    </div>
	 
     {{template "common/left.tpl" .}}
  </div>

     {{template "common/foot.tpl"}}