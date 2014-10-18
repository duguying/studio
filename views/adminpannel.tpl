<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		
		<link rel="stylesheet" href="/static/css/style.min.css">
		<link rel="stylesheet" href="/static/css/admin.min.css">
		<link rel="stylesheet" type="text/css" media="all" href="/static/syntaxhighlighter/styles/shCoreDefault.css" />
		<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>

		<script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
	    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
	    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>

	</head>
  	<body>
  		<div class="main">
			<div class="admin-left-banner">
				<div class="admin-title">独孤影</div>
				<div class="admin-menu-left">
					<ul>
						<li id="new-ariticle">新文章</li>
						<li id="article-manage">文章管理</li>
						<li id="attach-manage">附件管理</li>
						<li id="oss-manage">OSS管理</li>
					</ul>
				</div>
			</div>
			<div class="admin-right-pannel">
				<div class="admin-top-menu">
					<div id="default_menu"></div>
					<div id="new-article-menu"></div>
					<div id="article-manage-menu"></div>
					<div id="attach-manage-menu"></div>
					<div id="oss-manage-menu"></div>
				</div>
				<div class="admin-main-pannel" >

					<div id="default-box">
						<ul class="ds-recent-comments" data-num-items="10" data-show-avatars="1" data-show-time="1" data-show-admin="1" data-excerpt-length="70"></ul>
						<!--多说js加载开始，一个页面只需要加载一次 -->
						<script type="text/javascript">
						var duoshuoQuery = {short_name:"duguying"};
						(function() {
						    var ds = document.createElement('script');
						    ds.type = 'text/javascript';ds.async = true;
						    ds.src = 'http://static.duoshuo.com/embed.js';
						    ds.charset = 'UTF-8';
						    (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(ds);
						})();
						</script>
						<!--多说js加载结束，一个页面只需要加载一次 -->
					</div>

					<div id="new-article-box"></div>
					<div id="article-manage-box"></div>
					<div id="attach-manage-box">
						<input id="img" type="file" size="45" name="attach" class="input">
						<button class="button" id="buttonUpload" onclick="alert('hello');">上传</button>
					</div>
					<div id="oss-manage-box"></div>
				</div>
			</div>
		</div>
		<script src="/static/js/admin.footer.min.js"></script>
	</body>
</html>

