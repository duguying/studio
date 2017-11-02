<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>reset password</title>
</head>
<body>
	<hr>
	<div>
		<h3>用户{{{.username}}}: 重置密码</h3>
		<form action="/password/reset" method="post">
			<input type="password" name="password" id="">
			<input type="submit" value="修改">
		</form>
	</div>
</body>
</html>