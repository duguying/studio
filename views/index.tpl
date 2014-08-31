<!DOCTYPE html>

<html>
  	<head>
    	<title>blog</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.2.0/css/bootstrap.min.css">
		<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.2.0/css/bootstrap-theme.min.css">
		<link rel="stylesheet" href="/static/css/style.css">
		<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
		<script src="http://cdn.bootcss.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
	</head>
  	<body>
  		<div class="main">
  			<div class="header">
  				<div class="icons">
  					<span class="icon-github icon"></span>
  					<span class="icon-weibo icon"></span>
  					<span class="icon-gplus icon"></span>
  					<span class="icon-twitter icon"></span>
  				</div>
  				<ul class="menu">
  					<li>关于</li>
  					<li>项目</li>
  					<li>列表</li>
  					<li>博文</li>
  				</ul>
  				<div class="banner">
  					<span class="title">独孤影的博客</span>
  				</div>
  				<div class="gap">
  					<span class="gravatar"></span>
  				</div>
  			</div>
			<div class="article_list">
				{{range $k,$v := .articles_in_page}}
					<div class="article">
						<h2>{{$v.title}}</h2>
						<div>{{str2html $v.content}}</div>
					</div>
					<hr>
				{{end}}
				{{if .prev_page_flag}}
				<a href="/page/{{.prev_page}}">上一页</a>
				{{end}}
				{{if .next_page_flag}}
				<a href="/page/{{.next_page}}">下一页</a>
				{{end}}
			</div>
			<div class="box_right">
				<h4>按月统计</h4>
				{{range $k,$v := .count_by_month}}
					<span>{{$v.month}}[{{$v.number}}]</span><br>
				{{end}}
			</div>
			<div class="footer">
				<div class="links">
					<ul>
						<li>link1</li>
						<li>link2</li>
						<li>link3</li>
						<li>link4</li>
						<li>link5</li>
					</ul>
				</div>
				<div class="copyright">©2014 the theme designed by Rex Lee inspired by <a href="https://www.byvoid.com/">byvoid</a>, the program written by Rex Lee with Golang base on <a href="http://beego.me/">Beego</a> framework.</div>
			</div>
		</div>
	</body>
</html>
