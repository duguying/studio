<!DOCTYPE html>

<html>
  	<head>
    	<title>独孤影</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		
		<link rel="stylesheet" href="/static/css/style.css">
		<link rel="stylesheet" href="/static/css/admin.css">
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
					<div id="new-article-menu"></div>
					<div id="article-manage-menu"></div>
					<div id="attach-manage-menu"></div>
					<div id="oss-manage-menu"></div>
				</div>
				<div class="admin-main-pannel" >
					<div id="new-article-box"></div>


					<div id="article-manage-box">
						<ul>
							<li class="admin-amb-banner">
								<ul>
									<li class="admin-col-id admin-amb-cell">ID</li>
									<li class="admin-col-title admin-amb-cell">文章标题</li>
									<li class="admin-col-author admin-amb-cell">作者</li>
									<li class="admin-col-time admin-amb-cell">发布时间</li>
									<li class="admin-col-count admin-amb-cell">阅读量</li>
									<li class="admin-col-operation admin-amb-cell">操作</li>
								</ul>
							</li>
							<li class="admin-amb-line">
								<ul>
									<li class="admin-col-id admin-amb-cell">100</li>
									<li class="admin-col-title admin-amb-cell">这是一个很长很长很长很长很长很长很长很长的标题</li>
									<li class="admin-col-author admin-amb-cell">内容</li>
									<li class="admin-col-time admin-amb-cell">内容</li>
									<li class="admin-col-count admin-amb-cell">内容</li>
									<li class="admin-col-operation admin-amb-cell"><a href="" style="padding-right: 5px;">辑</a><a href="" style="padding-left: 5px;">删</a></li>
								</ul>
							</li>
							
							<li class="admin-amb-banner">
								<ul>
									<li class="admin-col-id admin-amb-cell">ID</li>
									<li class="admin-col-title admin-amb-cell">文章标题</li>
									<li class="admin-col-author admin-amb-cell">作者</li>
									<li class="admin-col-time admin-amb-cell">发布时间</li>
									<li class="admin-col-count admin-amb-cell">阅读量</li>
									<li class="admin-col-operation admin-amb-cell">操作</li>
								</ul>
							</li>
						</ul>
					</div>


					<div id="attach-manage-box"></div>
					<div id="oss-manage-box"></div>
				</div>
			</div>
		</div>
		<script src="/static/js/admin.js"></script>
	</body>
</html>

