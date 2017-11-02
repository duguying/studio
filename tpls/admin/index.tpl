<!DOCTYPE html>

<html ng-app="Admin">
	<head>
		<title>独孤影 - {{global.title}}</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		{{{if .inDev}}}
				{{{template "inc/css_dev.tpl" .}}}
		{{{else}}}
				{{{template "inc/css_prod.tpl" .}}}
		{{{end}}}
	</head>
  	<body>
  		<div class="admin">
  			<div class="left" ng-controller="NavsController">
	  			<ul>
	  				<li ng-class="{active: global.currentPath==''}">
	  					<a href="/admin/" title="首页"><i class="icon iconfont">&#xe616;</i></a>
	  				</li>
	  				<li ng-class="{active: global.currentPath=='new_article'}">
	  					<a href="/admin/new_article" title="新建文章"><i class="icon iconfont">&#xe6cf;</i></a>
	  				</li>
	  				<li ng-class="{active: global.currentPath=='manage_article'}">
	  					<a href="/admin/manage_article" title="文章管理"><i class="icon iconfont">&#xe701;</i></a>
	  				</li>
	  				<li ng-class="{active: global.currentPath=='manage_project'}">
	  					<a href="/admin/manage_project" title="项目管理"><i class="icon iconfont">&#xe604;</i></a>
	  				</li>
	  				<li ng-class="{active: global.currentPath=='manage_oss'}">
	  					<a href="/admin/manage_oss" title="OSS管理"><i class="icon iconfont">&#xe78e;</i></a>
	  				</li>
	  			</ul>
  			</div>
  			<div class="right" ng-view></div>
  		</div>
	</body>
	{{{if .inDev}}}
			{{{template "../inc/js_dev.tpl" .}}}
	{{{else}}}
			{{{template "../inc/js_prod.tpl" .}}}
	{{{end}}}
</html>
