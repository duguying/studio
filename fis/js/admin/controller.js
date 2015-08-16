"use strict";

function NavsController ($scope, $http, $location) {
	$http.get("/api/admin/navlist",null).success(function (data) {
		$scope.navs = data;
		// $scope.isActive = true;
		$scope.currentPath = $location.path().replace("/admin","").replace("/","")
	});
}

function IndexController($scope,$rootScope){
	// $scope.active = true
    console.log("hello index");
    $rootScope.global = {
    	title: "首页",
    	currentPath: ""
    }
    
}

function NewArticleController($scope,$rootScope){
    console.log("hello new article");
    $rootScope.global = {
    	title: "添加文章",
    	currentPath: "new_article"
    }
}

function ManageArticleController($scope,$rootScope){
    console.log("hello manage article");
    $rootScope.global = {
    	title: "管理文章",
    	currentPath: "manage_article"
    }
}

function ManageProjectController($scope,$rootScope){
    console.log("hello manage project");
    $rootScope.global = {
    	title: "管理项目",
    	currentPath: "manage_project"
    }
}

function ManageOssController($scope,$rootScope){
    console.log("hello manage oss");
    $rootScope.global = {
    	title: "OSS管理",
    	currentPath: "manage_oss"
    }
}


