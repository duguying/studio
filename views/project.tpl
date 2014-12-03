<!DOCTYPE html>

<html>
	<head>
		<title>独孤影 - 项目</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<meta content="独孤影,博客,项目" name="keywords" />
		<meta content="独孤影的项目" name="description" />
		<link rel="stylesheet" href="/static/css/style.min.css">
		<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
		<script src="/static/js/jquery.min.js"></script>
	</head>
	<body >
		<div class="main">

  			{{template "inc/header.tpl" .}}

			<div class="projects">
				<h1><a href="/project">项目</a></h1>
				<div class="intro">
					此处包含了本人接触编程以来所开发或参与开发的项目，这些项目多为开源项目，欢迎参阅。
				</div>
				{{range $k,$v := .projects_in_page}}
					<div class="project-box">
						<img src="{{$v.icon_url}}" alt="">
						<div class="info">
							<span class="name">{{$v.name}}</span><hr>
							<span class="author">作者: {{$v.author}}</span>
							<span class="time">时间: {{$v.time}}</span>
							<div>{{$v.description}}</div>
						</div>
					</div>
					
				{{end}}
				{{if .prev_page_flag}}
				<a href="{{.prev_page}}" class="page-nav">上一页</a>
				{{end}}
				{{if .next_page_flag}}
				<a href="{{.next_page}}" class="page-nav">下一页</a>
				{{end}}
			</div>

			{{template "inc/footer.tpl" .}}

		</div>
		
	</body>
</html>

