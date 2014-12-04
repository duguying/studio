<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影 - {{.title}}</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<meta content="独孤影,博客,{{.keywords}}" name="keywords" />
		<meta content="{{.title}}" name="description" />
		{{asset "sass/style.scss"}}
		{{asset "syntaxhighlighter/styles/shCoreDefault.css"}}
	</head>
  	<body >
  		<div class="main">

  			{{template "inc/header.tpl" .}}

  			<div class="article-list">
	  			<div class="article" itemscope itemtype="http://schema.org/Article">
	  				<a class="article-title" title="{{.title}}" href="/article/{{.uri}}" itemprop="name">{{.title}}</a>
	  				<div class="article-ps">
						Tag {{.keywords|tags|str2html}} on <span datetime="{{.time}}" itemprop="datePublished" class="post-time">{{.time}}</span> by <span title="作者: {{.author}}" itemprop="author" class="author-name">{{.author}}</span> view <span title="{{.count}}次阅读" class="view-count">{{.count}}</span>
					</div>
					<div class="article-content" itemprop="articleBody">
						{{str2html .content}}
					</div>

					<!-- 多说评论框 start -->
					<div class="ds-thread" data-thread-key="{{.id}}" data-title="{{.title}}" data-url="http://duguying.net/article/{{.title}}" id="comments"></div>
					<!-- 多说评论框 end -->
					<!-- 多说公共JS代码 start (一个网页只需插入一次) -->
					<script type="text/javascript">
						var duoshuoQuery = {short_name:"{{.duoshuo}}"};
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

			{{asset "js/article.js"}}
			
		</div>
	</body>
</html>
