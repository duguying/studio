<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影 - {{.title}}</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<meta content="独孤影,博客,{{.keywords}}" name="keywords" />
		<meta content="{{.title}}" name="description" />
		<link rel="stylesheet" href="/static/src/bin/css/style.min.css">
		<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
		<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
	</head>
  	<body style='display: none'>
  		<div class="main">

  			{{template "inc/header.tpl" .}}

  			<div class="article-list">
	  			<div class="article" itemscope itemtype="http://schema.org/Article">
	  				<a class="article-title" title="{{.title}}" href="/article/{{.uri}}" itemprop="name">{{.title}}</a>
	  				<div class="article-ps">
						Tag {{.keywords|tags|str2html}} on <a datetime="{{.time}}" itemprop="datePublished">{{.time}}</a> by <a title="作者: {{.author}}" itemprop="author">{{.author}}</a> view <a title="{{.count}}次阅读">{{.count}}</a>
					</div>
					<div class="article-content" itemprop="articleBody">
						{{str2html .content}}
					</div>

					<!-- 多说评论框 start -->
					<div class="ds-thread" data-thread-key="{{.id}}" data-title="{{.title}}" data-url="http://duguying.net/article/{{.title}}"></div>
					<!-- 多说评论框 end -->
					<!-- 多说公共JS代码 start (一个网页只需插入一次) -->
					<script type="text/javascript">
						var duoshuoQuery = {short_name:"duguying"};
						(function() {
							var ds = document.createElement('script');
							ds.type = 'text/javascript';ds.async = true;
							ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
							ds.charset = 'UTF-8';
							(document.getElementsByTagName('head')[0] 
							 || document.getElementsByTagName('body')[0]).appendChild(ds);
						})();
					</script>
					<!-- 多说公共JS代码 end -->

				</div>
			</div>

			{{template "inc/rightbar.tpl" .}}

			{{template "inc/footer.tpl" .}}

		</div>
	</body>
</html>
