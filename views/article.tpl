<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Article Page {{.id}}</title>
</head>
<body>
	<div>
		<h1>{{.title}}</h1>
		<span>{{.time}}</span><br>
		<span>Author {{.author}}</span><br>
		<span>Count:{{.count}}</span><br>
		<span>Keywords {{.keywords}}</span><br><br>
		<div>{{.content}}</div>
	</div>
</body>
</html>