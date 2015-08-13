"use strict";

function NavsController ($scope, $http, $location) {
	$http.get("/api/admin/navlist",null).success(function (data) {
		$scope.navs = data;
		// $scope.isActive = true;
		$scope.currentPath = $location.path().replace("/admin","").replace("/","")
	});
}

function IndexController($scope){
	// $scope.active = true
    console.log("hello index");
    
}

function NewArticleController($scope){
    console.log("hello new article");
}

function ManageArticleController($scope){
    console.log("hello manage article");
}

function ManageProjectController($scope){
    console.log("hello manage project");
}

function ManageOssController($scope){
    console.log("hello manage oss");
}


