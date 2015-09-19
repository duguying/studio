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
  		<div class="main">
  			<div class="left" ng-controller="NavsController">
	  			<ul>
	  				<li ng-repeat="nav in navs" ng-class="{active: global.currentPath=='{{nav.uri}}'}">
	  					<a href="/admin/{{nav.uri}}">{{nav.title}}</a>
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
