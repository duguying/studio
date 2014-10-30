<!DOCTYPE html>

<html>
	<head>
		<title>独孤影 {{.title}}</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta content="独孤影,博客,个人网站,IT,技术,编程" name="keywords" />
		<meta content="独孤影的博客" name="description" />
		<link rel="EditURI" type="application/rsd+xml" title="RSD" href="{{.host}}/xmlrpc" />
		<link rel="stylesheet" href="/static/css/style.min.css">
		<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
		<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
		<meta name="google-site-verification" content="ohMjRPHv0sKAahvl1H0GC7Dx0-z-zXbMNnWBfxp2PYY" />
	</head>
	<body >
		<div class="main">

  			{{template "inc/header.tpl" .}}

			<div class="article-list">
				{{range $k,$v := .articles_in_page}}
					<div class="article" itemscope itemtype="http://schema.org/Article">
						<a class="article-title" title="{{$v.title}}" href="/article/{{$v.uri}}" itemprop="name">{{$v.title}}</a>
						<div class="article-ps">
							Tag {{$v.keywords|tags|str2html}} on <span datetime="{{$v.time}}" itemprop="datePublished" class="post-time">{{$v.time}}</span> by <span title="作者: {{$v.author}}" itemprop="author" class="author-name">{{$v.author}}</span> view <span title="{{$v.count}}次阅读" class="view-count">{{$v.count}}</span>
						</div>
						<div class="article-content" itemprop="articleBody">
							{{str2html $v.content}}
						</div>
					</div>
					<hr>
				{{end}}
				{{if .prev_page_flag}}
				<a href="{{.prev_page}}" class="page-nav">上一页</a>
				{{end}}
				{{if .next_page_flag}}
				<a href="{{.next_page}}" class="page-nav">下一页</a>
				{{end}}
			</div>

			{{template "inc/rightbar.tpl" .}}

			{{template "inc/footer.tpl" .}}

		</div>
		
	</body>
</html>

