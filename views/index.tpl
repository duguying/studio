<!DOCTYPE html>

<html>
  	<head>
    	<title>This is just a test website for Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
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
	</body>
</html>
