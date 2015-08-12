<!DOCTYPE html>

<html ng-app>
	<head>
		<title>独孤影</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		{{{asset "sass/style.scss"}}}
		{{{asset "sass/admin.scss"}}}
		{{{asset "syntaxhighlighter/styles/shCoreDefault.css"}}}
		{{{asset "js/global/angular.min.js"}}}
		{{{asset "js/global/angular-route.min.js"}}}
		{{{asset "js/admin/admin.js"}}}
	</head>
  	<body ng-controller="Admin.TestController">
		<div class="main">
			<div class="left">
				<ul>
					<li ng-repeat="nav in navs" uri="{{nav.uri}}">
						{{nav.title}}
					</li>
				</ul>
			</div>
			<div class="right">
				<div class="top">
					;
				</div>
				<div class="pannel">
					
				</div>
			</div>
		</div>
		<div class="footer">
			
		</div>
	</body>
</html>
