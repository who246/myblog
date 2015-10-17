CCFBLOG(Golang development)
==== 
基于beego框架Golang语言开发的博客
演示地址：http://blog.lcsg.pw/
#首页
![ccfblog](http://i13.tietuku.com/6eedcc2e716004f0.png) 
#内容页
![ccfblog](http://i13.tietuku.com/33c4c65163b2eea9.jpg) 
-------
#Install

##Installation
`go get github.com/who246/myblog`
##config
modify /conf/app.conf
##build
`go build`
##run 
`./myblog`

#LOG

##2015-08-10
 1.完成页面的展示的功能
##2015-08-20
 2.完成后台操作的功能
##2015-09-26
1.增加自动更新并重启的功能(不构建)
##2015-09-27
1.增加重新构建并重启的功能(需要配置环境)
##2015-10-12
1.增加redis文章缓存和点击量缓存的功能
2.增加自动建表的功能