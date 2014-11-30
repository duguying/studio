<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影 - 找回密码</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="stylesheet" href="/static/src/bin/css/style.min.css">
		<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
		<script src="/static/js/jquery.min.js"></script>
		<style>
		.center{
			display: block;
		    margin: 0 auto;
		    position: relative;
		    width: 300px;
		    text-align: center;
		}
		.center .input{
			width: 100%;
		}
		</style>
	</head>
  	<body>
  		<div class="main">
  			<div class="center">
				<label for="">Getback Password</label><br>
				<input type="text" name="username" id="" class="input" placeholder="usernmae"><br>
				<button>找回</button>
			</div>
			<div class="footer">
				<div class="copyright">©2014 the theme designed by Rex Lee inspired by <a href="https://www.byvoid.com/">byvoid</a>, the program written by Rex Lee with Golang base on <a href="http://beego.me/">Beego</a> framework.</div>
			</div>
 		</div>
	</body>
	<script>
	$(document).ready(function (e) {
		$("button").click(function (e) {
			var user_name = $("input[name='username']").val();
			$.ajax({
				type: "get",
				url: "/password/sendemail",
				data: { username: user_name },
				dataType: "json",
				success: function(msg){
					console.log(msg);
					if (msg.result) {
						alert("邮件发送成功-"+msg.msg);
					} else{
						alert("邮件发送失败-"+msg.msg);
					};
				}
			});
		});
		
	})
	</script>
</html>
