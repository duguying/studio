<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="stylesheet" href="/static/css/style.css">
		<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
		<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
	</head>
  	<body>
  		<div class="main">

  			{{template "inc/header.tpl" .}}

			<div class="article-list">
				{{range $k,$v := .articles_in_page}}
					<div class="article">
						<a class="article-title" title="{{$v.title}}" href="/article/{{$v.uri}}">{{$v.title}}</a>
						<div class="article-ps">
							Tag {{$v.keywords|tags|str2html}} on <a>{{$v.time}}</a> by <a title="作者: {{$v.author}}">{{$v.author}}</a> view <a title="{{$v.count}}次阅读">{{$v.count}}</a>
						</div>
						<div class="article-content">
							{{str2html $v.content}}
						</div>
					</div>
					<hr>
				{{end}}
				{{if .prev_page_flag}}
				<a href="/page/{{.prev_page}}" class="page-nav">上一页</a>
				{{end}}
				{{if .next_page_flag}}
				<a href="/page/{{.next_page}}" class="page-nav">下一页</a>
				{{end}}
			</div>


			{{template "inc/rightbar.tpl" .}}

			{{template "inc/footer.tpl" .}}


		</div>
	</body>
</html>

