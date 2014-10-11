<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>独孤影 - 文章列表</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta content="独孤影,博客,个人网站,IT,技术,编程" name="keywords" />
	<meta content="博文列表" name="description" />
	<link rel="stylesheet" href="/static/src/bin/css/style.min.css">
	<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
	<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
</head>
<body >
	<div class="main">
  		{{template "inc/header.tpl" .}}

		<div class="lp-article-list">
			<ul>
				{{range $k,$v := .list}}
				<li class="a-l-item">
					<ul>
						<li class="time">{{$v.time}}</li>
						<li class="title"><a href="/article/{{$v.uri}}">{{$v.title}}</a></li>
					</ul>
				</li>
				{{end}}
			</ul>
		</div>
		<div class="lp-nav">
			{{if .prev_page_flag}}
			<a href="/list/{{.prev_page}}" class="page-nav">上一页</a>
			{{end}}
			{{if .next_page_flag}}
			<a href="/list/{{.next_page}}" class="page-nav">下一页</a>
			{{end}}
		</div>

  		{{template "inc/footer.tpl" .}}
  	</div>
</body>
</html>