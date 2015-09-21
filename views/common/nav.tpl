<body>

    <div class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <button class="navbar-toggle collapsed" type="button" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand hidden-sm" href="/" onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'navbar-首页'])">CCFBLOG</a>
        </div>
        <div class="navbar-collapse collapse" role="navigation">
          <ul class="nav navbar-nav">
            <li ><a href="/"    onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'v2doc'])">首页</a></li>
			<li ><a href="/login"    onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'v2doc'])">管理</a></li>
<!--           <li><a href="http://v3.bootcss.com/" target="_blank" onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'v3doc'])">Bootstrap3中文文档</a></li>
            <li><a href="http://v4.bootcss.com/" target="_blank" onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'v4doc'])">Bootstrap 4.0 预览</a></li>
            <li><a href="/p/lesscss/" target="_blank" onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'less'])">Less 教程</a></li>
            <li><a href="http://www.jquery123.com/" target="_blank" onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'jquery'])">jQuery API</a></li>
 -->           
           
          </ul>
         
          <ul class="nav navbar-nav navbar-right hidden-sm">
		 <li>
            	 <form class="navbar-form navbar-left" role="/search">
						<div class="form-group">
							<input type="text" name="key" class="form-control" value="{{.key}}" placeholder="输入搜索内容" />
						</div> 
						
					</form>
            </li>
            <li><a href="/about/" onclick="_hmt.push(['_trackEvent', 'navbar', 'click', 'about'])">关于</a></li>
          </ul>
        </div>
      </div>
    </div>