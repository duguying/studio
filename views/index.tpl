<!DOCTYPE html>

<html>
  	<head>
    	<title>blog</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.2.0/css/bootstrap.min.css">
		<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.2.0/css/bootstrap-theme.min.css">
		<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
		<script src="http://cdn.bootcss.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
	</head>
  	<body>
		<div>
			<h3>月份文章统计列表</h3>
			{{range $k,$v := .count_by_month}}
				<span>{{$v.month}}[{{$v.number}}]</span><br>
			{{end}}
		</div>
		<hr>
		<div>
			<h3>当前页文章列表</h3>
			{{range $k,$v := .articles_in_page}}
				<div style="border:1px solid gray;">{{$v.title}}<br>{{$v.content}}</div><br>
			{{end}}
			{{if .prev_page_flag}}
			<a href="/?page={{.prev_page}}">上一页</a>
			{{end}}
			{{if .next_page_flag}}
			<a href="/?page={{.next_page}}">下一页</a>
			{{end}}
		</div>
	</body>
</html>
