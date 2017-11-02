<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>registor</title>
	<link rel="stylesheet" href="/static/css/style.min.css">
</head>
<body>
	<div class="main">
		<form action="/registor" method="post" class="center registor">
			<span>用户注册</span>
			<span><label for="">用户名</label><input type="text" name="username" id=""></span>
			<span><label for="">密码</label><input type="password" name="password" id=""></span>
			<input type="submit" value="注册">
		</form>
		<div class="footer">
			<div class="copyright">©2014 the theme designed by Rex Lee, the program written by Rex Lee with Golang base on <a href="http://beego.me/">Beego</a> framework.</div>
		</div>
	</div>
</body>
</html>