<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		{{{asset "sass/style.scss"}}}
		{{{asset "sass/admin.scss"}}}
		{{{asset "syntaxhighlighter/styles/shCoreDefault.css"}}}
		{{{asset "js/global/jquery.min.js"}}}

		<script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
	    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
	    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>

	</head>
  	<body>
  		<div class="admin">
  			<div class="main-header">
  				<ul>
  					<li><a target="_blank" href="/logout">登出</a></li>
  					<li><a target="_blank" href="/">首页</a></li>
  				</ul>
  			</div>
			<div class="admin-left-banner">
				<div class="admin-title">独孤影</div>
				<div class="admin-menu-left">
					<ul>
						<li id="new-article">新文章</li>
						<li id="article-manage">文章管理</li>
						<li id="attach-manage">附件管理</li>
						<li id="oss-manage">OSS管理</li>
						<li id="project-manage">项目管理</li>
					</ul>
				</div>
			</div>
			<div class="admin-right-pannel">
				<div class="admin-top-menu">
					<div id="default_menu"></div>
					<div id="new-article-menu">
						<label for="article-title" style="margin-left: 10px;color:white;">文章标题</label>
						<input type="text" name="title" id="article-title" style="margin-left: 10px;margin-top: 7px;width: 250px;">
						<label for="article-tags" style="color: white;margin-left: 10px;margin-right: 10px;">关键词</label>
						<input type="text" name="tags" placeholder="逗号,分隔" id="article-tags" class="article-tags">
						<button id="submit" style="margin-left: 10px;background-color: white;border-radius: 4px;">发布</button>
					</div>
					<div id="article-manage-menu"></div>
					<div id="attach-manage-menu"></div>
					<div id="oss-manage-menu"></div>
					<div id="project-manage-menu"></div>
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

					<div id="editor-box">
						<textarea class="m-input d-input" name="content" id="myEditor" style="width:100%;height:430px;"></textarea>
					</div>
					<div id="article-manage-box">
						<div class="loading"></div>
						<div class="list"></div>
					</div>
					<div id="attach-manage-box">
						<input id="img" type="file" size="45" name="attach" class="input">
						<button class="button" id="buttonUpload" onclick="alert('hello');">上传</button>
					</div>
					<div id="oss-manage-box"></div>
					<div id="project-manage-box"></div>
				</div>
			</div>
		</div>
		{{{asset "js/admin.js"}}}
	</body>
</html>

