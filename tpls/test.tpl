<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>tests</title>
</head>
<body>
	<div>
			<h3>添加文章</h3>
			<form action="/add" method="post">
				<label for="">标题</label><input type="text" name="title" id=""><br>
				<label for="">关键字</label><input type="text" name="keywords" id=""><br>
				<label for="">正文</label><br><textarea name="content" id="" cols="30" rows="10"></textarea><br>
				<input type="submit" value="添加文章">
			</form>
		</div>
		<hr>
		<div>
			<h3>修改文章 id = 1</h3>
			<form action="/update" method="post">
				<input type="hidden" name="id" value="1">
				<label for="">标题</label><input type="text" name="title" id=""><br>
				<label for="">关键字</label><input type="text" name="keywords" id=""><br>
				<label for="">正文</label><br><textarea name="content" id="" cols="30" rows="10"></textarea><br>
				<input type="submit" value="修改文章">
			</form>
		</div>
		<hr>
		<div>
			<h3>删除文章</h3>
			<form action="/delete" method="post">
				<label for="">id</label><input type="text" name="id" id=""><br>
				<input type="submit" value="删除文章">
			</form>
		</div>
		<hr>
		<div>
			<h3>修改用户名，当前用户[{{{.username}}}]</h3>
			<form action="/rename" method="post">
				<label for="">新用户名</label><input type="text" name="username" id=""><br>
				<input type="submit" value="修改">
			</form>
		</div>
		<hr>
		<div>
			<h3>修改Email</h3>
			<form action="/email" method="post">
				<label for="">Email</label><input type="text" name="email" id=""><br>
				<input type="submit" value="修改">
			</form>
		</div>
		<hr>
		<div>
			<h3>发送Email验证找回密码</h3>
			<form action="/password/sendemail" method="get">
				<input type="submit" value="发送">
			</form>
		</div>
		<hr>
		<div>
			<h3>修改密码</h3>
			<form action="/password/change" method="post">
				<label for="">旧密码</label><input type="password" name="oldpassword" id="">
				<label for="">新密码</label><input type="password" name="password" id="">
				<input type="submit" value="发送">
			</form>
		</div>
		<hr>
		<div>
			<h3>上传文件</h3>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="file" id="">
				<input type="submit" value="上传">
			</form>
		</div>
		<hr>
		<div>
			<h3>修改密码</h3>
			<form action="/password/change" method="post">
				<label for="">输入密码</label><input type="password" name="old_password"><br>
				<label for="">新密码</label><input type="password" name="password"><br>
				<input type="submit" value="修改">
			</form>
		</div>
</body>
</html>