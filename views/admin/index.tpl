<!DOCTYPE html>

<html ng-app="Admin">
	<head>
		<title>独孤影 - {{global.title}}</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		{{{asset "sass/style.scss"}}}
		{{{asset "sass/admin.scss"}}}
		{{{asset "js/global/angular.min.js"}}}
		{{{asset "js/global/angular-route.min.js"}}}
		{{{asset "js/admin/admin.js"}}}
		{{{asset "js/admin/controller.js"}}}
		{{{asset "js/admin/directive.js"}}}
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
</html>
