<!DOCTYPE html>

<html ng-app="Admin">
	<head>
		<title>独孤影 - {{global.title}}</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<link rel="stylesheet" href="/static/css/style.css">
		<link rel="stylesheet" href="/static/css/admin.css">
		<script src="/static/js/global/angular.min.js"></script>
		<script src="/static/js/global/angular-route.min.js"></script>
		<script src="/static/ueditor/ueditor.config.js"></script>
		<script src="/static/ueditor/ueditor.all.js"></script>
		<script src="/static/ueditor/angular-ueditor.js"></script>
		<script src="/static/js/admin/admin.js"></script>
		<script src="/static/js/admin/directive.js"></script>
		<script src="/static/js/admin/controller.js"></script>
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
