
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="X-UA-Compatible" content="IE=Edge">
<meta charset="utf-8">

<title>bootstrap-wysihtml5</title>


<link rel="stylesheet" type="text/css" href="http://jhollingworth.github.io/bootstrap-wysihtml5/lib/css/bootstrap.min.css"></link>
 
<link rel="stylesheet" type="text/css" href="http://jhollingworth.github.io/bootstrap-wysihtml5/src/bootstrap-wysihtml5.css"></link>

<style type="text/css" media="screen">
  .btn.jumbo {
    font-size: 20px;
    font-weight: normal;
    padding: 14px 24px;
    margin-right: 10px;
    -webkit-border-radius: 6px;
    -moz-border-radius: 6px;
    border-radius: 6px;
  }
</style>

 
</head>
<body >
<div class="container">
  <div class="hero-unit" style="margin-top:40px">
    <h1 style="font-size:58px">bootstrap-wysihtml5 <small>Simple, beautiful wysiwyg editors</small></h1>
    <hr/>
	<form method="POST" action="/article/add">
	
	{{range .err}}
	<div style="color:#F00">
	{{.Key}}{{.Message}}
     </div>
	{{end}}
	<pre><code>zxczxc</code></pre>
	<label for="title">类型</label>
    <input type="title" class="form-control" id="typeId" name="typeId" placeholder="类型" value="{{.typeId}}"/>
	<label for="title">标题</label>
    <input type="title" class="form-control" id="title" name="title" placeholder="标题" value="{{.title}}"/>
    <textarea id="some-textarea"name="content" class="textarea" placeholder="Enter text ..." style="width: 810px; height: 200px">{{.content}}</textarea>
	<input class="btn btn-danger" type="submit" value="Submit">
	</form>
  </div>
   
</div>


<script src="http://jhollingworth.github.io/bootstrap-wysihtml5/lib/js/wysihtml5-0.3.0.js"></script>
<script src="http://jhollingworth.github.io/bootstrap-wysihtml5/lib/js/jquery-1.7.2.min.js"></script>
 
<script src="http://jhollingworth.github.io/bootstrap-wysihtml5/lib/js/bootstrap.min.js"></script>
<script src="/static/js/bootstrap-wysihtml5.js"></script>
<link rel="stylesheet" href="http://yandex.st/highlightjs/8.0/styles/solarized_dark.min.css">
<script src="http://yandex.st/highlightjs/8.0/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>

<script>
$('#some-textarea').wysihtml5({
	"font-styles": true, //Font styling, e.g. h1, h2, etc. Default true
	"emphasis": true, //Italics, bold, etc. Default true
	"lists": true, //(Un)ordered lists, e.g. Bullets, Numbers. Default true
	"html": true, //Button which allows you to edit the generated HTML. Default false
	"link": true, //Button to insert a link. Default true
	"image": true, //Button to insert an image. Default true,
	 "code":true, 
	"color": true //Button to change color of font  
});
 
</script>

 

</body>
</html>
